# null0go

A WASM-based game engine for Go, where games ("carts") are loaded as zip files containing `main.wasm`.

Part of the [null0](https://github.com/notnullgames/null0) project.

I am playing with a few go-ideas here:

- host made in go. I dunno if I will finish this, but want to play with the idea of using ebiten/wazero/merged_fs, since it's simpler than the C counterparts
- carts made with tinygo. This idea will probably stick around.
- codegen for host & carts, written in go

## Quick Start

```sh
# Install dependencies (requires go & tinygo)
make setup

# Build the host runtime
make null0

# Build example carts
make carts

# Run a cart
./host/null0 hello.null0
```

The host merges carts and directories into a single filesystem, so you can test a directory with main.wasm, or give a cart access to any directories you want:

```sh
./host/null0 hello.null0 someassets/
```

this makes it so `main.wasm` in `hello.null0` has access to all the files in `someassets/`, as if they were in the cart.

## Usage

Here is a quick usage-example (go cart)

```go
import null0 "github.com/notnullgames/null0go/cartapi"

func main() {
  red := null0.Color{R: 255, G: 0, B: 0, A: 255}
  null0.Clear(red)

  if null0.KeyDown(null0.KEY_SPACE) {
      println("Space pressed!")
  }

  points := []null0.Vector{
      {X: 10, Y: 10},
      {X: 50, Y: 50},
      {X: 10, Y: 50},
  }

  null0.DrawPolygon(red, 3, points)
  img := null0.LoadImage("player.png")
}
```

You can also export [callbacks](https://notnull.games/null0/cart).

## Code Generation

The entire API is defined in YAML files under `api/`. Run the code generator with:

```bash
make codegen
```

> [!WARNING]
> This stubs the host-code, and will wipe it!

This generates:

1. **`cartapi/null0.go`** - Go package with types, enums, and `//go:wasmimport null0 <func>` declarations
2. **`host/null0_api.go`** - Host-side function stubs (currently just prints debug messages)

The generator converts snake_case function names to PascalCase for Go exports (e.g., `delta_time` → `DeltaTime`), while preserving the original names in the wasmimport directives (so it will work with carts already made in other languages.)
