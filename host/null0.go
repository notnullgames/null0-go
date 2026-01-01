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
    "github.com/tetratelabs/wazero/sys"
    "github.com/urfave/cli/v3"
    "github.com/yalue/merged_fs"
    "github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
    ctx context.Context
    mod api.Module
    cbUpdate api.Function
    cbButtonUp api.Function
    cbButtonDown api.Function
    cbKeyUp api.Function
    cbKeyDown api.Function
    cbMouseDown api.Function
    cbMouseUp api.Function
    cbMouseMoved api.Function

    // Track input state for edge detection
    prevKeys map[ebiten.Key]bool
    prevMouseButtons map[ebiten.MouseButton]bool
    prevMouseX, prevMouseY int
}

func (game *Game) Update() error {
    // Initialize tracking maps on first run
    if game.prevKeys == nil {
        game.prevKeys = make(map[ebiten.Key]bool)
        game.prevMouseButtons = make(map[ebiten.MouseButton]bool)
        game.prevMouseX, game.prevMouseY = ebiten.CursorPosition()
    }

    // Track all keys for edge detection
    allKeys := []ebiten.Key{
        ebiten.KeySpace, ebiten.KeyApostrophe, ebiten.KeyComma, ebiten.KeyMinus, ebiten.KeyPeriod, ebiten.KeySlash,
        ebiten.Key0, ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9,
        ebiten.KeySemicolon, ebiten.KeyEqual,
        ebiten.KeyA, ebiten.KeyB, ebiten.KeyC, ebiten.KeyD, ebiten.KeyE, ebiten.KeyF, ebiten.KeyG, ebiten.KeyH, ebiten.KeyI, ebiten.KeyJ,
        ebiten.KeyK, ebiten.KeyL, ebiten.KeyM, ebiten.KeyN, ebiten.KeyO, ebiten.KeyP, ebiten.KeyQ, ebiten.KeyR, ebiten.KeyS, ebiten.KeyT,
        ebiten.KeyU, ebiten.KeyV, ebiten.KeyW, ebiten.KeyX, ebiten.KeyY, ebiten.KeyZ,
        ebiten.KeyLeftBracket, ebiten.KeyBackslash, ebiten.KeyRightBracket, ebiten.KeyGraveAccent,
        ebiten.KeyEscape, ebiten.KeyEnter, ebiten.KeyTab, ebiten.KeyBackspace, ebiten.KeyInsert, ebiten.KeyDelete,
        ebiten.KeyRight, ebiten.KeyLeft, ebiten.KeyDown, ebiten.KeyUp,
        ebiten.KeyPageUp, ebiten.KeyPageDown, ebiten.KeyHome, ebiten.KeyEnd,
        ebiten.KeyCapsLock, ebiten.KeyScrollLock, ebiten.KeyNumLock, ebiten.KeyPrintScreen, ebiten.KeyPause,
        ebiten.KeyF1, ebiten.KeyF2, ebiten.KeyF3, ebiten.KeyF4, ebiten.KeyF5, ebiten.KeyF6,
        ebiten.KeyF7, ebiten.KeyF8, ebiten.KeyF9, ebiten.KeyF10, ebiten.KeyF11, ebiten.KeyF12,
        ebiten.KeyShiftLeft, ebiten.KeyControlLeft, ebiten.KeyAltLeft, ebiten.KeyMetaLeft,
        ebiten.KeyShiftRight, ebiten.KeyControlRight, ebiten.KeyAltRight, ebiten.KeyMetaRight,
        ebiten.KeyContextMenu,
    }

    // Check for key events
    if game.cbKeyDown != nil || game.cbKeyUp != nil {
        for _, key := range allKeys {
            pressed := ebiten.IsKeyPressed(key)
            wasPressed := game.prevKeys[key]

            if pressed && !wasPressed && game.cbKeyDown != nil {
                // Key down event
                game.cbKeyDown.Call(game.ctx, uint64(key))
            } else if !pressed && wasPressed && game.cbKeyUp != nil {
                // Key up event
                game.cbKeyUp.Call(game.ctx, uint64(key))
            }

            game.prevKeys[key] = pressed
        }
    }

    // Check for mouse button events
    mouseButtons := []ebiten.MouseButton{ebiten.MouseButtonLeft, ebiten.MouseButtonRight, ebiten.MouseButtonMiddle}
    if game.cbMouseDown != nil || game.cbMouseUp != nil {
        for _, button := range mouseButtons {
            pressed := ebiten.IsMouseButtonPressed(button)
            wasPressed := game.prevMouseButtons[button]

            if pressed && !wasPressed && game.cbMouseDown != nil {
                // Mouse down event
                game.cbMouseDown.Call(game.ctx, uint64(button))
            } else if !pressed && wasPressed && game.cbMouseUp != nil {
                // Mouse up event
                game.cbMouseUp.Call(game.ctx, uint64(button))
            }

            game.prevMouseButtons[button] = pressed
        }
    }

    // Check for mouse movement
    if game.cbMouseMoved != nil {
        mouseX, mouseY := ebiten.CursorPosition()
        if mouseX != game.prevMouseX || mouseY != game.prevMouseY {
            game.cbMouseMoved.Call(game.ctx, api.EncodeF32(float32(mouseX)), api.EncodeF32(float32(mouseY)))
            game.prevMouseX, game.prevMouseY = mouseX, mouseY
        }
    }

    return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
    if game.cbUpdate != nil {
        game.cbUpdate.Call(game.ctx)
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 640, 480
}

func main() {
    app := &cli.Command{
        Name:  "null0",
        Usage: "run a cart from one or more dirs/zips",
        Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:    "net",
                Aliases: []string{"n"},
                Usage:   "Enable cart networking",
            },
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

// runCart loads main.wasm from fsys, wires it into a WASI runtime and calls init function
func runCart(name string, fsys fs.FS, enableNet bool) error {
    game := &Game{}
    ctx := context.Background()
    game.ctx = ctx

    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle(name)

    wasmBytes, err := readFileFromFS(fsys, "main.wasm")
    if err != nil {
        return fmt.Errorf("cart %q: main.wasm not found or unreadable: %w", name, err)
    }

    r := wazero.NewRuntime(ctx)
    defer r.Close(ctx)

    fsConfig := wazero.NewFSConfig().WithFSMount(fsys, "/")

    // Create persistent WASI with custom proc_exit and fd_write
    wasiBuilder := r.NewHostModuleBuilder("wasi_snapshot_preview1")
    registerPersistentWASI(wasiBuilder, fsys)
    if _, err := wasiBuilder.Instantiate(ctx); err != nil {
        return fmt.Errorf("instantiate custom WASI: %w", err)
    }

    // Register null0 API
    null0Builder := r.NewHostModuleBuilder("null0")
    registerNull0API(null0Builder)
    if _, err := null0Builder.Instantiate(ctx); err != nil {
        return fmt.Errorf("instantiate null0 module: %w", err)
    }

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

    // Store the module for callback access
    game.mod = mod

    // Call main() directly instead of _start to avoid WASI cleanup
    mainFn := mod.ExportedFunction("main")
    if mainFn != nil {
        if _, err := mainFn.Call(ctx); err != nil {
            return fmt.Errorf("cart %q: main failed: %w", name, err)
        }
    } else {
        // Fall back to _start if main doesn't exist
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

    // Get callback functions before starting the game loop
    game.cbUpdate = mod.ExportedFunction("update")
    game.cbButtonUp = mod.ExportedFunction("buttonUp")
    game.cbButtonDown = mod.ExportedFunction("buttonDown")
    game.cbKeyUp = mod.ExportedFunction("keyUp")
    game.cbKeyDown = mod.ExportedFunction("keyDown")
    game.cbMouseDown = mod.ExportedFunction("mouseDown")
    game.cbMouseUp = mod.ExportedFunction("mouseUp")
    game.cbMouseMoved = mod.ExportedFunction("mouseMoved")

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }

    // Test callback functions (in a real implementation, these would be called by the game loop)
    // buttonUp(button: GamepadButton, player: u32)
    // if fn := mod.ExportedFunction("buttonUp"); fn != nil {
    //     fn.Call(ctx, uint64(0), uint64(0)) // button=0, player=0
    // }

    // // buttonDown(button: GamepadButton, player: u32)
    // if fn := mod.ExportedFunction("buttonDown"); fn != nil {
    //     fn.Call(ctx, uint64(1), uint64(0)) // button=1, player=0
    // }

    // // keyUp(key: Key)
    // if fn := mod.ExportedFunction("keyUp"); fn != nil {
    //     fn.Call(ctx, uint64(65)) // key=65 (example: 'A')
    // }

    // // keyDown(key: Key)
    // if fn := mod.ExportedFunction("keyDown"); fn != nil {
    //     fn.Call(ctx, uint64(66)) // key=66 (example: 'B')
    // }

    // // mouseDown(button: MouseButton)
    // if fn := mod.ExportedFunction("mouseDown"); fn != nil {
    //     fn.Call(ctx, uint64(0)) // button=0 (left button)
    // }

    // // mouseUp(button: MouseButton)
    // if fn := mod.ExportedFunction("mouseUp"); fn != nil {
    //     fn.Call(ctx, uint64(0)) // button=0 (left button)
    // }

    // // mouseMoved(x: f32, y: f32)
    // if fn := mod.ExportedFunction("mouseMoved"); fn != nil {
    //     fn.Call(ctx, api.EncodeF32(100.5), api.EncodeF32(200.5)) // x=100.5, y=200.5
    // }

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
