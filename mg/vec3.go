package mg

type Vec3 struct {
	X Float
	Y Float
	Z Float
}

func V3(x, y, z Float) Vec3 {
	return Vec3{x, y, z}
}

func V3i(x, y, z int) Vec3 {
	return Vec3{Float(x), Float(y), Float(z)}
}

func Rgb(r, g, b Float) Vec3 {
	return V3(r, g, b)
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vec3) Mul(b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vec3) Div(b Vec3) Vec3 {
	return Vec3{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func (a Vec3) Negate() Vec3 {
	return Vec3{-a.X, -a.Y, -a.Z}
}

func (a Vec3) Min(b Vec3) Vec3 {
	return Vec3{min(a.X, b.X), min(a.Y, b.Y), min(a.Z, b.Z)}
}

func (a Vec3) Max(b Vec3) Vec3 {
	return Vec3{max(a.X, b.X), max(a.Y, b.Y), max(a.Z, b.Z)}
}

func (a Vec3) Clamp(min, max Vec3) Vec3 {
	return Vec3{
		X: Clamp(a.X, min.X, max.X),
		Y: Clamp(a.Y, min.Y, max.Y),
		Z: Clamp(a.Z, min.Z, max.Z),
	}
}

func (a Vec3) Lerp(b Vec3, t Float) Vec3 {
	return Vec3{
		X: Lerp(a.X, b.X, t),
		Y: Lerp(a.Y, b.Y, t),
		Z: Lerp(a.Z, b.Z, t),
	}
}

func (a Vec3) CloseEnough(b Vec3) bool {
	return CloseEnough(a.X, b.X) && CloseEnough(a.Y, b.Y) && CloseEnough(a.Z, b.Z)
}

func (a Vec3) Invert() Vec3 {
	return Vec3{1 / a.X, 1 / a.Y, 1 / a.Z}
}

func (a Vec3) Mag() Float {
	return Sqrt(a.MagSqr())
}

func (a Vec3) MagSqr() Float {
	return a.Dot(a)
}

func (a Vec3) Dot(b Vec3) Float {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (a Vec3) Normalize() Vec3 {
	return a.Mulf(InvSqrt(a.Dot(a)))
}

func (a Vec3) Distance(b Vec3) Float {
	return Sqrt(a.DistanceSqr(b))
}

func (a Vec3) DistanceSqr(b Vec3) Float {
	dx := b.X - a.X
	dy := b.Y - a.Y
	dz := b.Z - a.Z
	return dx*dx + dy*dy + dz*dz
}

func (a Vec3) AngleBetween(b Vec3) Angle {
	return FromRad(Atan2(a.Cross(b).Mag(), a.Dot(b)))
}

// @todo(judah): Rotate

func (a Vec3) Reflect(normal Vec3) Vec3 {
	dot := a.Dot(normal)
	return Vec3{
		a.X - (2*normal.X)*dot,
		a.Y - (2*normal.Y)*dot,
		a.Z - (2*normal.Z)*dot,
	}
}

func (a Vec3) Components() (x, y, z Float) {
	return a.X, a.Y, a.Z
}

func (a Vec3) Addf(f Float) Vec3 { return a.Add(Vec3{f, f, f}) }
func (a Vec3) Subf(f Float) Vec3 { return a.Sub(Vec3{f, f, f}) }
func (a Vec3) Mulf(f Float) Vec3 { return a.Mul(Vec3{f, f, f}) }
func (a Vec3) Divf(f Float) Vec3 { return a.Div(Vec3{f, f, f}) }

func (a Vec3) Clampf(min, max Float) Vec3 {
	return a.Clamp(Vec3{min, min, min}, Vec3{max, max, max})
}

// Implements the [color.Color] interface
func (v Vec3) RGBA() (r, g, b, a uint32) {
	r = uint32(v.X * 255)
	r |= r << 8

	g = uint32(v.Y * 255)
	g |= g << 8

	b = uint32(v.Z * 255)
	b |= b << 8

	return
}
