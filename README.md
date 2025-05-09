# jc.go

Packages for various things I usually want when writing Go.

## What's included

### Array

General purpose array types.

- `stable.go` - Stable dynamic array. Pointers taken will not invalidate while resizing.

### MG (Math for Games)

There are many (game focused) math packages in Go, but I usually don't like them or forget what they're called, so I end up rolling my own for each project. This package is so I no longer have to do that.

- `rect.go` - Rectangle operations
   - `rectcut.go` - Simple implementation of RectCut
- `float.go` - Floating-point math operations
   - Build with `MG_USE_F32` to use `float32` internally
   - Build with `MG_USE_F64` to use `float64` internally (default)
- `angle.go` - Utilities for working with different angle units
   - Build with `MG_USE_RADIANS` to use radians internally
   - Build with `MG_USE_DEGREES` to use degrees internally
   - Build with `MG_USE_TURNS` to use turns internally (default)
- `v2.go` - Vector2 operations
- `v3.go` - Vector3 operations
- `v4.go` - Vector4 operations

### Thelp (Test Helpers)

Tiny utils I want to make tests nicer to write.

## LICENSE

See [LICENSE](./LICENSE)
