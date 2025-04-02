package gm

type Vec2 struct {
	X Float
	Y Float
}

func V2(x, y Float) Vec2 {
	return Vec2{x, y}
}

func V2i(x, y int) Vec2 {
	return Vec2{Float(x), Float(y)}
}

func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

func (a Vec2) Mul(b Vec2) Vec2 {
	return Vec2{a.X * b.X, a.Y * b.Y}
}

func (a Vec2) Div(b Vec2) Vec2 {
	return Vec2{a.X / b.X, a.Y / b.Y}
}

func (a Vec2) Negate() Vec2 {
	return Vec2{-a.X, -a.Y}
}

func (a Vec2) Min(b Vec2) Vec2 {
	return Vec2{min(a.X, b.X), min(a.Y, b.Y)}
}

func (a Vec2) Max(b Vec2) Vec2 {
	return Vec2{max(a.X, b.X), max(a.Y, b.Y)}
}

func (a Vec2) Clamp(min, max Vec2) Vec2 {
	return Vec2{
		X: Clamp(a.X, min.X, max.X),
		Y: Clamp(a.Y, min.Y, max.Y),
	}
}

func (a Vec2) Lerp(b Vec2, t Float) Vec2 {
	return Vec2{
		X: Lerp(a.X, b.X, t),
		Y: Lerp(a.Y, b.Y, t),
	}
}

func (a Vec2) CloseEnough(b Vec2) bool {
	return CloseEnough(a.X, b.X) && CloseEnough(a.Y, b.Y)
}

func (a Vec2) Invert() Vec2 {
	return Vec2{1 / a.X, 1 / a.Y}
}

func (a Vec2) Mag() Float {
	return Sqrt(a.MagSqr())
}

func (a Vec2) MagSqr() Float {
	return a.Dot(a)
}

func (a Vec2) Dot(b Vec2) Float {
	return a.X*b.X + a.Y*b.Y
}

func (a Vec2) Normalize() Vec2 {
	return a.Mulf(InvSqrt(a.Dot(a)))
}

func (a Vec2) Distance(b Vec2) Float {
	return Sqrt(a.DistanceSqr(b))
}

func (a Vec2) DistanceSqr(b Vec2) Float {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}

func (a Vec2) AngleBetween(b Vec2) Angle {
	return FromRad(Atan2(b.Y, b.X) - Atan2(a.Y, a.X))
}

func (a Vec2) Rotate(angle Angle) Vec2 {
	sin := Float(angle.Sin())
	cos := Float(angle.Cos())
	return Vec2{a.X*cos - a.Y*sin, a.X*sin + a.Y*cos}
}

func (a Vec2) Reflect(normal Vec2) Vec2 {
	dot := a.Dot(normal)
	return Vec2{a.X - (2*normal.X)*dot, a.Y - (2*normal.Y)*dot}
}

func (a Vec2) Components() (x, y Float) {
	return a.X, a.Y
}

func (a Vec2) Addf(f Float) Vec2 { return a.Add(Vec2{f, f}) }
func (a Vec2) Subf(f Float) Vec2 { return a.Sub(Vec2{f, f}) }
func (a Vec2) Mulf(f Float) Vec2 { return a.Mul(Vec2{f, f}) }
func (a Vec2) Divf(f Float) Vec2 { return a.Div(Vec2{f, f}) }

func (a Vec2) Clampf(min, max Float) Vec2 {
	return a.Clamp(Vec2{min, min}, Vec2{max, max})
}
