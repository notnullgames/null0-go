package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type TypeDef struct {
	Description    string            `yaml:"description"`
	CArg           string            `yaml:"c_arg"`
	CRet           string            `yaml:"c_ret"`
	Host           string            `yaml:"host"`
	RequiresLength bool              `yaml:"requiresLength"`
	Members        map[string]string `yaml:"members"`
	Enums          map[string]int    `yaml:"enums"`
}

type FuncArg struct {
	Type string
	Name string
}

type FuncDef struct {
	Description string            `yaml:"description"`
	Args        map[string]string `yaml:"args"`
	Returns     string            `yaml:"returns"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Determine base directory (parent of codegen)
	baseDir := ".."
	if _, err := os.Stat("api"); err == nil {
		// Already in the right directory
		baseDir = "."
	}

	// Load type definitions
	types, err := loadTypes(filepath.Join(baseDir, "api/types.yml"))
	if err != nil {
		return fmt.Errorf("load types: %w", err)
	}

	// Load all API files
	apiFiles := []string{
		"api/graphics.yml",
		"api/input.yml",
		"api/sound.yml",
		"api/utilities.yml",
	}

	allFuncs := make(map[string]FuncDef)
	for _, file := range apiFiles {
		funcs, err := loadFunctions(filepath.Join(baseDir, file))
		if err != nil {
			return fmt.Errorf("load %s: %w", file, err)
		}
		for name, def := range funcs {
			allFuncs[name] = def
		}
	}

	// Load callbacks
	callbacks, err := loadFunctions(filepath.Join(baseDir, "api/callbacks.yml"))
	if err != nil {
		return fmt.Errorf("load callbacks: %w", err)
	}

	// Generate cart header
	if err := generateCartHeader(types, allFuncs, filepath.Join(baseDir, "cartapi/null0.go")); err != nil {
		return fmt.Errorf("generate cart header: %w", err)
	}

	// Generate host stubs
	if err := generateHostStubs(types, allFuncs, callbacks, filepath.Join(baseDir, "host/null0_api.go")); err != nil {
		return fmt.Errorf("generate host stubs: %w", err)
	}

	fmt.Println("✓ Generated cartapi/null0.go")
	fmt.Println("✓ Generated host/null0_api.go")
	return nil
}

func loadTypes(filename string) (map[string]TypeDef, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	types := make(map[string]TypeDef)
	if err := yaml.Unmarshal(data, &types); err != nil {
		return nil, err
	}

	return types, nil
}

func loadFunctions(filename string) (map[string]FuncDef, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	funcs := make(map[string]FuncDef)
	if err := yaml.Unmarshal(data, &funcs); err != nil {
		return nil, err
	}

	return funcs, nil
}

func yamlTypeToGo(yamlType string, types map[string]TypeDef) string {
	// Handle arrays
	if strings.HasSuffix(yamlType, "[]") {
		return "uint32" // arrays are pointers in WASM
	}

	// Check if it's a defined type with members (struct) - pass as pointer
	if typeDef, ok := types[yamlType]; ok {
		if len(typeDef.Members) > 0 {
			// Struct - pass as pointer in WASM ABI
			return "uint32"
		}
		if len(typeDef.Enums) > 0 {
			// Enum - pass as value
			return "uint32"
		}
	}

	// Map basic types
	switch yamlType {
	case "void":
		return ""
	case "i8":
		return "int8"
	case "i16":
		return "int16"
	case "i32":
		return "int32"
	case "i64":
		return "int64"
	case "u8":
		return "uint8"
	case "u16":
		return "uint16"
	case "u32":
		return "uint32"
	case "u64":
		return "uint64"
	case "f32":
		return "float32"
	case "f64":
		return "float64"
	case "bool":
		return "bool"
	case "string":
		return "uint32" // pointer in WASM
	case "Image", "Font", "Sound":
		return "uint32"
	default:
		return "uint32" // unknown types default to pointer
	}
}

func snakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

func sanitizeParamName(name string) string {
	// List of Go reserved keywords
	keywords := map[string]bool{
		"break": true, "case": true, "chan": true, "const": true, "continue": true,
		"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
		"func": true, "go": true, "goto": true, "if": true, "import": true,
		"interface": true, "map": true, "package": true, "range": true, "return": true,
		"select": true, "struct": true, "switch": true, "type": true, "var": true,
	}

	if keywords[name] {
		return name + "_"
	}
	return name
}

func yamlTypeToWasm(yamlType string, types map[string]TypeDef) string {
	if strings.HasSuffix(yamlType, "[]") {
		return "i32" // pointer
	}

	switch yamlType {
	case "i8", "i16", "i32", "u8", "u16", "u32", "bool":
		return "i32"
	case "i64", "u64":
		return "i64"
	case "f32":
		return "f32"
	case "f64":
		return "f64"
	case "string":
		return "i32"
	case "Image", "Font", "Sound":
		return "i32"
	case "void":
		return ""
	default:
		// Structs/enums - check if it's a type with members (struct) or enum
		if typeDef, ok := types[yamlType]; ok {
			if len(typeDef.Members) > 0 {
				// Struct - passed as pointer
				return "i32"
			}
			if len(typeDef.Enums) > 0 {
				// Enum - passed as value
				return "i32"
			}
		}
		return "i32"
	}
}

func generateCartHeader(types map[string]TypeDef, funcs map[string]FuncDef, outputFile string) error {
	var out strings.Builder

	out.WriteString("package null0\n\n")
	out.WriteString("// This file is auto-generated from api/*.yml - DO NOT EDIT\n\n")

	// Generate type definitions
	out.WriteString("// Type Definitions\n\n")

	// Collect all type names and sort them
	var typeNames []string
	for name := range types {
		if name != "void" && (len(types[name].Members) > 0 || len(types[name].Enums) > 0) {
			typeNames = append(typeNames, name)
		}
	}
	sort.Strings(typeNames)

	// Generate structs
	for _, name := range typeNames {
		typeDef := types[name]
		if len(typeDef.Members) > 0 {
			out.WriteString(fmt.Sprintf("// %s\n", typeDef.Description))
			out.WriteString(fmt.Sprintf("type %s struct {\n", name))
			for memberName, memberType := range typeDef.Members {
				goType := yamlTypeToGo(memberType, types)
				// Capitalize first letter for export
				exportedName := strings.ToUpper(memberName[:1]) + memberName[1:]
				out.WriteString(fmt.Sprintf("\t%s %s\n", exportedName, goType))
			}
			out.WriteString("}\n\n")
		}
	}

	// Generate enums
	for _, name := range typeNames {
		typeDef := types[name]
		if len(typeDef.Enums) > 0 {
			out.WriteString(fmt.Sprintf("// %s\n", typeDef.Description))
			out.WriteString(fmt.Sprintf("type %s uint32\n\n", name))
			out.WriteString("const (\n")

			// Sort enum values
			var enumNames []string
			for enumName := range typeDef.Enums {
				enumNames = append(enumNames, enumName)
			}
			sort.Strings(enumNames)

			for _, enumName := range enumNames {
				enumValue := typeDef.Enums[enumName]
				out.WriteString(fmt.Sprintf("\t%s %s = %d\n", enumName, name, enumValue))
			}
			out.WriteString(")\n\n")
		}
	}

	// Generate function imports
	out.WriteString("// API Functions\n\n")

	// Sort function names
	var funcNames []string
	for name := range funcs {
		funcNames = append(funcNames, name)
	}
	sort.Strings(funcNames)

	for _, name := range funcNames {
		funcDef := funcs[name]

		// Generate comment
		out.WriteString(fmt.Sprintf("// %s\n", funcDef.Description))

		// Build parameter list
		var params []string
		var paramNames []string
		if funcDef.Args != nil {
			for argName := range funcDef.Args {
				paramNames = append(paramNames, argName)
			}
			sort.Strings(paramNames)

			for _, argName := range paramNames {
				argType := funcDef.Args[argName]
				goType := yamlTypeToGo(argType, types)
				safeArgName := sanitizeParamName(argName)
				params = append(params, fmt.Sprintf("%s %s", safeArgName, goType))
			}
		}

		// Build return type
		returnType := ""
		if funcDef.Returns != "" && funcDef.Returns != "void" {
			returnType = " " + yamlTypeToGo(funcDef.Returns, types)
		}

		// Convert function name to PascalCase for Go
		goFuncName := snakeToPascal(name)

		// Generate the wasmimport directive with original name
		out.WriteString(fmt.Sprintf("//go:wasmimport null0 %s\n", name))
		out.WriteString(fmt.Sprintf("func %s(%s)%s\n\n", goFuncName, strings.Join(params, ", "), returnType))
	}

	// Write to file
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		return err
	}
	return os.WriteFile(outputFile, []byte(out.String()), 0644)
}

func generateHostStubs(types map[string]TypeDef, funcs map[string]FuncDef, callbacks map[string]FuncDef, outputFile string) error {
	var out strings.Builder

	out.WriteString("package main\n\n")
	out.WriteString("// This file is auto-generated from api/*.yml - DO NOT EDIT\n\n")
	out.WriteString("import (\n")
	out.WriteString("\t\"context\"\n")
	out.WriteString("\t\"fmt\"\n\n")
	out.WriteString("\t\"github.com/tetratelabs/wazero\"\n")
	out.WriteString("\t\"github.com/tetratelabs/wazero/api\"\n")
	out.WriteString(")\n\n")

	out.WriteString("// registerNull0API registers all null0 API functions with the host module builder.\n")
	out.WriteString("func registerNull0API(builder wazero.HostModuleBuilder) {\n")

	// Sort function names
	var funcNames []string
	for name := range funcs {
		funcNames = append(funcNames, name)
	}
	sort.Strings(funcNames)

	for _, name := range funcNames {
		funcDef := funcs[name]

		// Build parameter types for wazero
		var paramTypes []string
		var paramNames []string
		if funcDef.Args != nil {
			for argName := range funcDef.Args {
				paramNames = append(paramNames, argName)
			}
			sort.Strings(paramNames)

			for _, argName := range paramNames {
				argType := funcDef.Args[argName]
				wasmType := yamlTypeToWasm(argType, types)
				paramTypes = append(paramTypes, fmt.Sprintf("api.ValueType%s", strings.ToUpper(wasmType[:1])+wasmType[1:]))
			}
		}

		// Build return types
		var returnTypes []string
		if funcDef.Returns != "" && funcDef.Returns != "void" {
			wasmType := yamlTypeToWasm(funcDef.Returns, types)
			returnTypes = append(returnTypes, fmt.Sprintf("api.ValueType%s", strings.ToUpper(wasmType[:1])+wasmType[1:]))
		}

		returnTypesStr := "nil"
		if len(returnTypes) > 0 {
			returnTypesStr = "[]api.ValueType{" + strings.Join(returnTypes, ", ") + "}"
		}

		paramTypesStr := "nil"
		if len(paramTypes) > 0 {
			paramTypesStr = "[]api.ValueType{" + strings.Join(paramTypes, ", ") + "}"
		}

		out.WriteString(fmt.Sprintf("\t// %s: %s\n", name, funcDef.Description))
		out.WriteString("\tbuilder.NewFunctionBuilder().\n")
		out.WriteString("\t\tWithGoModuleFunction(\n")
		out.WriteString("\t\t\tapi.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {\n")
		out.WriteString(fmt.Sprintf("\t\t\t\tfmt.Println(\"[null0] %s called (stub)\")\n", name))
		out.WriteString("\t\t\t}),\n")
		out.WriteString(fmt.Sprintf("\t\t\t%s,\n", paramTypesStr))
		out.WriteString(fmt.Sprintf("\t\t\t%s,\n", returnTypesStr))
		out.WriteString("\t\t).\n")
		out.WriteString(fmt.Sprintf("\t\tExport(\"%s\")\n\n", name))
	}

	out.WriteString("}\n")

	// Write to file
	return os.WriteFile(outputFile, []byte(out.String()), 0644)
}
