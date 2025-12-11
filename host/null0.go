package main

import (
    "archive/zip"
    "bytes"
    "context"
    "fmt"
    "io"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/tetratelabs/wazero"
    "github.com/tetratelabs/wazero/api"
    "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
    "github.com/tetratelabs/wazero/sys"
    "github.com/urfave/cli/v3"
    "github.com/yalue/merged_fs"
)

func main() {
    app := &cli.Command{
        Name:  "null0",
        Usage: "run a cart from one or more dirs/zips",
        Flags: []cli.Flag{
            // &cli.BoolFlag{
            //     Name:    "net",
            //     Aliases: []string{"n"},
            //     Usage:   "Enable cart networking",
            // },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {
            if cmd.Args().Len() < 1 {
                cli.ShowRootCommandHelp(cmd)
                return cli.Exit("\nAt least one cart path is required.", 1)
            }

            firstPath := cmd.Args().Get(0)
            base := filepath.Base(firstPath)
            ext := filepath.Ext(base)
            name := strings.TrimSuffix(base, ext)

            netEnabled := cmd.Bool("net")

            var filesystems []fs.FS

            for i := 0; i < cmd.Args().Len(); i++ {
                p := cmd.Args().Get(i)

                info, err := os.Stat(p)
                if err != nil {
                    return err
                }

                if info.Mode().IsRegular() {
                    zr, err := openZipFS(p)
                    if err != nil {
                        return err
                    }
                    filesystems = append(filesystems, zr)
                } else if info.IsDir() {
                    filesystems = append(filesystems, os.DirFS(p))
                } else {
                    return cli.Exit("unsupported cart path type: "+p, 1)
                }
            }

            merged := merged_fs.MergeMultiple(filesystems...)

            return runCart(name, merged, netEnabled)
        },
    }

    if err := app.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}

// openZipFS opens a zip file at path and returns an fs.FS view of it.
func openZipFS(path string) (fs.FS, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    info, err := f.Stat()
    if err != nil {
        _ = f.Close()
        return nil, err
    }
    zr, err := zip.NewReader(f, info.Size())
    if err != nil {
        _ = f.Close()
        return nil, err
    }
    // zr implements fs.FS
    return zr, nil
}

// runCart loads main.wasm from fsys, wires it into a WASI runtime, exposes
// a "trace(ptr,len)" function, and calls the WASI _start entrypoint.
func runCart(name string, fsys fs.FS, enableNet bool) error {
    ctx := context.Background()

    wasmBytes, err := readFileFromFS(fsys, "main.wasm")
    if err != nil {
        return fmt.Errorf("cart %q: main.wasm not found or unreadable: %w", name, err)
    }

    r := wazero.NewRuntime(ctx)
    defer r.Close(ctx)

    if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
        return fmt.Errorf("instantiate WASI: %w", err)
    }

    // Host module "host" with function:
    //   (func (import "host" "trace") (param i32 i32))
    hostBuilder := r.NewHostModuleBuilder("host")

    traceFn := func(ctx context.Context, mod api.Module, ptr, length uint32) {
        mem := mod.Memory()
        if mem == nil {
            fmt.Fprintln(os.Stderr, "[trace] no memory")
            return
        }
        buf, ok := mem.Read(ptr, length)
        if !ok {
            fmt.Fprintln(os.Stderr, "[trace] out of bounds")
            return
        }
        fmt.Println("[trace]", string(buf))
    }

    // First two params are ctx and module; the rest are wasm params.
    hostBuilder.
        NewFunctionBuilder().
        WithGoModuleFunction(
            api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
                ptr := uint32(stack[0])
                length := uint32(stack[1])
                traceFn(ctx, mod, ptr, length)
            }),
            []api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
            nil,
        ).
        Export("trace")

    if _, err := hostBuilder.Instantiate(ctx); err != nil {
        return fmt.Errorf("instantiate host module: %w", err)
    }

    fsConfig := wazero.NewFSConfig().WithFSMount(fsys, "/")

    modConfig := wazero.NewModuleConfig().
        WithName(name).
        WithFSConfig(fsConfig).
        WithStdout(os.Stdout).
        WithStderr(os.Stderr)

    mod, err := r.InstantiateWithConfig(ctx, wasmBytes, modConfig)
    if err != nil {
        return fmt.Errorf("instantiate cart module: %w", err)
    }
    defer mod.Close(ctx)

    // Call initialization functions
    // Prefer load() over _start() because _start exits and closes file descriptors
    load := mod.ExportedFunction("load")
    if load != nil {
        if _, err := load.Call(ctx); err != nil {
            // Exit code 0 is success, not an error
            if exitErr, ok := err.(*sys.ExitError); !ok || exitErr.ExitCode() != 0 {
                return fmt.Errorf("cart %q: load failed: %w", name, err)
            }
        }
    } else {
        // Fall back to _start if load doesn't exist (for simple carts that just run and exit)
        start := mod.ExportedFunction("_start")
        if start != nil {
            if _, err := start.Call(ctx); err != nil {
                // Exit code 0 is success, not an error
                if exitErr, ok := err.(*sys.ExitError); !ok || exitErr.ExitCode() != 0 {
                    return fmt.Errorf("cart %q: _start failed: %w", name, err)
                }
            }
        }
    }

    // Test callback functions (in a real implementation, these would be called by the game loop)
    // buttonUp(button: GamepadButton, player: u32)
    if fn := mod.ExportedFunction("buttonUp"); fn != nil {
        fn.Call(ctx, uint64(0), uint64(0)) // button=0, player=0
    }

    // buttonDown(button: GamepadButton, player: u32)
    if fn := mod.ExportedFunction("buttonDown"); fn != nil {
        fn.Call(ctx, uint64(1), uint64(0)) // button=1, player=0
    }

    // keyUp(key: Key)
    if fn := mod.ExportedFunction("keyUp"); fn != nil {
        fn.Call(ctx, uint64(65)) // key=65 (example: 'A')
    }

    // keyDown(key: Key)
    if fn := mod.ExportedFunction("keyDown"); fn != nil {
        fn.Call(ctx, uint64(66)) // key=66 (example: 'B')
    }

    // mouseDown(button: MouseButton)
    if fn := mod.ExportedFunction("mouseDown"); fn != nil {
        fn.Call(ctx, uint64(0)) // button=0 (left button)
    }

    // mouseUp(button: MouseButton)
    if fn := mod.ExportedFunction("mouseUp"); fn != nil {
        fn.Call(ctx, uint64(0)) // button=0 (left button)
    }

    // mouseMoved(x: f32, y: f32)
    if fn := mod.ExportedFunction("mouseMoved"); fn != nil {
        fn.Call(ctx, api.EncodeF32(100.5), api.EncodeF32(200.5)) // x=100.5, y=200.5
    }

    return nil
}

// readFileFromFS reads a file from an fs.FS into a byte slice.
func readFileFromFS(fsys fs.FS, path string) ([]byte, error) {
    f, err := fsys.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var buf bytes.Buffer
    if _, err := io.Copy(&buf, f); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}
