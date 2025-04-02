package mg

type Vec4 struct {
	X Float
	Y Float
	Z Float
	W Float
}

func V4(x, y, z, w Float) Vec4 {
	return Vec4{x, y, z, w}
}

func V4i(x, y, z, w int) Vec4 {
	return Vec4{Float(x), Float(y), Float(z), Float(w)}
}

func Rgba(r, g, b, a Float) Vec4 {
	return V4(r, g, b, a)
}

func (a Vec4) Add(b Vec4) Vec4 {
	return Vec4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

func (a Vec4) Sub(b Vec4) Vec4 {
	return Vec4{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

func (a Vec4) Mul(b Vec4) Vec4 {
	return Vec4{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W}
}

func (a Vec4) Div(b Vec4) Vec4 {
	return Vec4{a.X / b.X, a.Y / b.Y, a.Z / b.Z, a.W / b.W}
}

func (a Vec4) Negate() Vec4 {
	return Vec4{-a.X, -a.Y, -a.Z, -a.W}
}

func (a Vec4) Min(b Vec4) Vec4 {
	return Vec4{min(a.X, b.X), min(a.Y, b.Y), min(a.Z, b.Z), min(a.W, b.W)}
}

func (a Vec4) Max(b Vec4) Vec4 {
	return Vec4{max(a.X, b.X), max(a.Y, b.Y), max(a.Z, b.Z), max(a.W, b.W)}
}

func (a Vec4) Clamp(min, max Vec4) Vec4 {
	return Vec4{
		X: Clamp(a.X, min.X, max.X),
		Y: Clamp(a.Y, min.Y, max.Y),
		Z: Clamp(a.Z, min.Z, max.Z),
		W: Clamp(a.W, min.W, max.W),
	}
}

func (a Vec4) Lerp(b Vec4, t Float) Vec4 {
	return Vec4{
		X: Lerp(a.X, b.X, t),
		Y: Lerp(a.Y, b.Y, t),
		Z: Lerp(a.Z, b.Z, t),
		W: Lerp(a.W, b.W, t),
	}
}

func (a Vec4) CloseEnough(b Vec4) bool {
	return CloseEnough(a.X, b.X) && CloseEnough(a.Y, b.Y) && CloseEnough(a.Z, b.Z) && CloseEnough(a.W, b.W)
}

func (a Vec4) Invert() Vec4 {
	return Vec4{1 / a.X, 1 / a.Y, 1 / a.Z, 1 / a.W}
}

func (a Vec4) Components() (x, y, z, w Float) {
	return a.X, a.Y, a.Z, a.W
}

func (a Vec4) Addf(f Float) Vec4 { return a.Add(Vec4{f, f, f, f}) }
func (a Vec4) Subf(f Float) Vec4 { return a.Sub(Vec4{f, f, f, f}) }
func (a Vec4) Mulf(f Float) Vec4 { return a.Mul(Vec4{f, f, f, f}) }
func (a Vec4) Divf(f Float) Vec4 { return a.Div(Vec4{f, f, f, f}) }

func (a Vec4) Clampf(min, max Float) Vec4 {
	return a.Clamp(Vec4{min, min, min, min}, Vec4{max, max, max, max})
}

// Implements the [color.Color] interface
func (v Vec4) RGBA() (r, g, b, a uint32) {
	r = uint32(v.X * 255)
	r |= r << 8

	g = uint32(v.Y * 255)
	g |= g << 8

	b = uint32(v.Z * 255)
	b |= b << 8

	a = uint32(v.W * 255)
	a |= a << 8

	return
}
