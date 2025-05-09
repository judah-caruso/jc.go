package mg

import "fmt"

// V3 is a 3-dimensional vector type
type V3 [3]Float

func (v V3) AsV2() V2         { return V2{v[0], v[1]} }
func (v V3) AsV4() V4         { return V4{v[0], v[1], v[2], 1.0} }
func (v V3) AsV4w(w Float) V4 { return V4{v[0], v[1], v[2], w} }

func (v V3) X() Float      { return v[0] }
func (v V3) Y() Float      { return v[1] }
func (v V3) Z() Float      { return v[2] }
func (v V3) R() Float      { return v[0] }
func (v V3) G() Float      { return v[1] }
func (v V3) B() Float      { return v[2] }
func (v V3) Width() Float  { return v[0] }
func (v V3) Height() Float { return v[1] }
func (v V3) Depth() Float  { return v[2] }

func (v *V3) Xp() *Float      { return &v[0] }
func (v *V3) Yp() *Float      { return &v[1] }
func (v *V3) Zp() *Float      { return &v[2] }
func (v *V3) Rp() *Float      { return &v[0] }
func (v *V3) Gp() *Float      { return &v[1] }
func (v *V3) Bp() *Float      { return &v[2] }
func (v *V3) Widthp() *Float  { return &v[0] }
func (v *V3) Heightp() *Float { return &v[1] }
func (v *V3) Depthp() *Float  { return &v[2] }

func (v *V3) SetX(f Float)      { v[0] = f }
func (v *V3) SetY(f Float)      { v[1] = f }
func (v *V3) SetZ(f Float)      { v[2] = f }
func (v *V3) SetR(f Float)      { v[0] = f }
func (v *V3) SetG(f Float)      { v[1] = f }
func (v *V3) SetB(f Float)      { v[2] = f }
func (v *V3) SetWidth(f Float)  { v[0] = f }
func (v *V3) SetHeight(f Float) { v[1] = f }
func (v *V3) SetDepth(f Float)  { v[2] = f }

func (v V3) Add(o V3) V3 {
	return V3{v[0] + o[0], v[1] + o[1], v[2] + o[2]}
}

func (v V3) Addf(f Float) V3 {
	return V3{v[0] + f, v[1] + f, v[2] + f}
}

func (v *V3) AddMut(o V3) {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
}

func (v *V3) AddMutf(f Float) {
	v[0] += f
	v[1] += f
	v[2] += f
}

func (v V3) Sub(o V3) V3 {
	return V3{v[0] - o[0], v[1] - o[1], v[2] - o[2]}
}

func (v V3) Subf(f Float) V3 {
	return V3{v[0] - f, v[1] - f, v[2] - f}
}

func (v *V3) SubMut(o V3) {
	v[0] -= o[0]
	v[1] -= o[1]
	v[2] -= o[2]
}

func (v *V3) SubMutf(f Float) {
	v[0] -= f
	v[1] -= f
	v[2] -= f
}

func (v V3) Mul(o V3) V3 {
	return V3{v[0] * o[0], v[1] * o[1], v[2] * o[2]}
}

func (v V3) Mulf(f Float) V3 {
	return V3{v[0] * f, v[1] * f, v[2] * f}
}

func (v *V3) MulMut(o V3) {
	v[0] *= o[0]
	v[1] *= o[1]
	v[2] *= o[2]
}

func (v *V3) MulMutf(f Float) {
	v[0] *= f
	v[1] *= f
	v[2] *= f
}

func (v V3) Div(o V3) V3 {
	return V3{v[0] / o[0], v[1] / o[1], v[2] / o[2]}
}

func (v V3) Divf(f Float) V3 {
	return V3{v[0] / f, v[1] / f, v[2] / f}
}

func (v *V3) DivMut(o V3) {
	v[0] /= o[0]
	v[1] /= o[1]
	v[2] /= o[2]
}

func (v *V3) DivMutf(f Float) {
	v[0] /= f
	v[1] /= f
	v[2] /= f
}

func (v V3) Eq(o V3) bool {
	return Abs(v[0]-o[0]) < FloatMin && Abs(v[1]-o[1]) < FloatMin && Abs(v[2]-o[2]) < FloatMin
}

func (v V3) CloseEnough(o V3) bool {
	return CloseEnough(v[0], o[0]) && CloseEnough(v[1], o[1]) && CloseEnough(v[2], o[2])
}

func (v V3) Dot(o V3) Float {
	return v[0]*o[0] + v[1]*o[1] + v[2]*o[2]
}

func (v V3) Mag() Float {
	return Sqrt(v.Dot(v))
}

func (v V3) Distance(to V3) Float {
	return v.Sub(to).Mag()
}

func (v V3) Angle(to V3) Angle {
	return FromRad(Acos(v.Dot(to) / (v.Mag() * to.Mag())))
}

func (v V3) Abs() V3 {
	return V3{Abs(v[0]), Abs(v[1]), Abs(v[2])}
}

func (v V3) Normalize() V3 {
	return v.Mulf(1.0 / v.Mag())
}

func (v *V3) NormalizeMut() {
	v.MulMutf(1.0 / v.Mag())
}

func (v V3) Invert() V3 {
	return V3{-v[0], -v[1], -v[2]}
}

func (v *V3) InvertMut() {
	v[0] = -v[0]
	v[1] = -v[1]
	v[2] = -v[2]
}

func (v V3) Clamp(min Float, max Float) V3 {
	return V3{Clamp(v[0], min, max), Clamp(v[1], min, max), Clamp(v[2], min, max)}
}

func (v *V3) ClampMut(min Float, max Float) {
	v[0] = Clamp(v[0], min, max)
	v[1] = Clamp(v[1], min, max)
	v[2] = Clamp(v[2], min, max)
}

func (v V3) Lerp(to V3, t Float) V3 {
	return V3{Lerp(v[0], to[0], t), Lerp(v[1], to[1], t), Lerp(v[2], to[2], t)}
}

func (v *V3) LerpMut(to V3, t Float) {
	v[0] = Lerp(v[0], to[0], t)
	v[1] = Lerp(v[1], to[1], t)
	v[2] = Lerp(v[2], to[2], t)
}

func (v V3) Cross(u V3) V3 {
	return V3{
		v[1]*u[2] - v[2]*u[1],
		v[2]*u[0] - v[0]*u[2],
		v[0]*u[1] - v[1]*u[0],
	}
}

// Implements the [color.Color] interface
func (v V3) RGBA() (r, g, b, a uint32) {
	r = uint32(v[0] * 0xFFFF)
	g = uint32(v[1] * 0xFFFF)
	b = uint32(v[2] * 0xFFFF)
	a = 0xFFFF
	return
}

// Implements the [fmt.Stringer] interface
func (v V3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v[0], v[1], v[2])
}
