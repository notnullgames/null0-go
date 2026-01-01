package main

import (
	"context"
	"encoding/binary"
	"io"
	"io/fs"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// persistentFDs stores file descriptors that survive proc_exit
var (
	persistentStdout io.Writer = os.Stdout
	persistentStderr io.Writer = os.Stderr
	persistentStdin  io.Reader = os.Stdin
)

// registerPersistentWASI registers a custom WASI implementation that keeps stdio open
func registerPersistentWASI(builder wazero.HostModuleBuilder, fsys fs.FS) {
	// fd_write - persistent implementation that always works
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fd := uint32(stack[0])
				iovs := uint32(stack[1])
				iovsLen := uint32(stack[2])
				resultNwritten := uint32(stack[3])

				mem := mod.Memory()
				var writer io.Writer

				switch fd {
				case 1: // stdout
					writer = persistentStdout
				case 2: // stderr
					writer = persistentStderr
				default:
					// Unsupported fd, return error
					stack[0] = 8 // EBADF
					return
				}

				var nwritten uint32
				for i := uint32(0); i < iovsLen; i++ {
					iovPtr := iovs + (i * 8)
					buf, ok := mem.Read(iovPtr, 8)
					if !ok {
						stack[0] = 21 // EINVAL
						return
					}

					bufPtr := binary.LittleEndian.Uint32(buf[0:4])
					bufLen := binary.LittleEndian.Uint32(buf[4:8])

					data, ok := mem.Read(bufPtr, bufLen)
					if !ok {
						stack[0] = 21 // EINVAL
						return
					}

					n, err := writer.Write(data)
					if err != nil {
						stack[0] = 5 // EIO
						return
					}
					nwritten += uint32(n)
				}

				// Write nwritten to result pointer
				if !mem.WriteUint32Le(resultNwritten, nwritten) {
					stack[0] = 21 // EINVAL
					return
				}

				stack[0] = 0 // Success
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("fd_write")

	// proc_exit - custom implementation that doesn't close stdio
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				// exitCode := uint32(stack[0])
				// Don't actually exit or close anything - just return
				// The module will stay alive for callbacks
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("proc_exit")

	// fd_close - no-op for stdout/stderr
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fd := uint32(stack[0])
				// Don't close stdout/stderr
				if fd == 1 || fd == 2 {
					stack[0] = 0 // Success but don't actually close
					return
				}
				stack[0] = 0 // Success
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("fd_close")

	// Add other minimal WASI functions needed for TinyGo
	// environ_sizes_get
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				environCount := uint32(stack[0])
				environBufSize := uint32(stack[1])
				mem := mod.Memory()
				mem.WriteUint32Le(environCount, 0)
				mem.WriteUint32Le(environBufSize, 0)
				stack[0] = 0 // Success
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("environ_sizes_get")

	// environ_get
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				stack[0] = 0 // Success (no environment variables)
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("environ_get")

	// fd_prestat_get - minimal implementation
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				stack[0] = 8 // EBADF - no preopened directories for now
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("fd_prestat_get")

	// fd_prestat_dir_name
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				stack[0] = 8 // EBADF
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("fd_prestat_dir_name")

	// fd_fdstat_get - minimal implementation
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fd := uint32(stack[0])
				statPtr := uint32(stack[1])
				mem := mod.Memory()

				if fd <= 2 { // stdin, stdout, stderr
					// Write a minimal fdstat structure (24 bytes)
					buf := make([]byte, 24)
					buf[0] = 2 // filetype: character device
					binary.LittleEndian.PutUint16(buf[2:], 0) // flags
					// rights_base and rights_inheriting set to 0
					mem.Write(statPtr, buf)
					stack[0] = 0 // Success
				} else {
					stack[0] = 8 // EBADF
				}
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("fd_fdstat_get")

	// clock_time_get
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				// clockID := uint32(stack[0])
				// precision := uint64(stack[1])
				resultPtr := uint32(stack[2])
				mem := mod.Memory()

				// Return a simple timestamp (nanoseconds)
				mem.WriteUint64Le(resultPtr, 0)
				stack[0] = 0 // Success
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI64, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("clock_time_get")

	// random_get
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				bufPtr := uint32(stack[0])
				bufLen := uint32(stack[1])
				mem := mod.Memory()

				// Fill with zeros for now (a real implementation would use crypto/rand)
				buf := make([]byte, bufLen)
				mem.Write(bufPtr, buf)
				stack[0] = 0 // Success
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("random_get")
}
