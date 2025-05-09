package mg

import "fmt"

// V4 is a 4-dimensional vector type
type V4 [4]Float

func (v V4) X() Float      { return v[0] }
func (v V4) Y() Float      { return v[1] }
func (v V4) Z() Float      { return v[2] }
func (v V4) W() Float      { return v[3] }
func (v V4) R() Float      { return v[0] }
func (v V4) G() Float      { return v[1] }
func (v V4) B() Float      { return v[2] }
func (v V4) A() Float      { return v[3] }
func (v V4) Width() Float  { return v[2] }
func (v V4) Height() Float { return v[3] }

func (v *V4) Xp() *Float      { return &v[0] }
func (v *V4) Yp() *Float      { return &v[1] }
func (v *V4) Zp() *Float      { return &v[2] }
func (v *V4) Wp() *Float      { return &v[3] }
func (v *V4) Rp() *Float      { return &v[0] }
func (v *V4) Gp() *Float      { return &v[1] }
func (v *V4) Bp() *Float      { return &v[2] }
func (v *V4) Ap() *Float      { return &v[3] }
func (v *V4) Widthp() *Float  { return &v[2] }
func (v *V4) Heightp() *Float { return &v[3] }

func (v *V4) SetX(f Float)      { v[0] = f }
func (v *V4) SetY(f Float)      { v[1] = f }
func (v *V4) SetZ(f Float)      { v[2] = f }
func (v *V4) SetW(f Float)      { v[3] = f }
func (v *V4) SetR(f Float)      { v[0] = f }
func (v *V4) SetG(f Float)      { v[1] = f }
func (v *V4) SetB(f Float)      { v[2] = f }
func (v *V4) SetA(f Float)      { v[3] = f }
func (v *V4) SetWidth(f Float)  { v[2] = f }
func (v *V4) SetHeight(f Float) { v[3] = f }

func (v V4) Add(o V4) V4 {
	return V4{v[0] + o[0], v[1] + o[1], v[2] + o[2], v[3] + o[3]}
}

func (v V4) Addf(f Float) V4 {
	return V4{v[0] + f, v[1] + f, v[2] + f, v[3] + f}
}

func (v *V4) AddMut(o V4) {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]
	v[3] += o[3]
}

func (v *V4) AddMutf(f Float) {
	v[0] += f
	v[1] += f
	v[2] += f
	v[3] += f
}

func (v V4) Sub(o V4) V4 {
	return V4{v[0] - o[0], v[1] - o[1], v[2] - o[2], v[3] - o[3]}
}

func (v V4) Subf(f Float) V4 {
	return V4{v[0] - f, v[1] - f, v[2] - f, v[3] - f}
}

func (v *V4) SubMut(o V4) {
	v[0] -= o[0]
	v[1] -= o[1]
	v[2] -= o[2]
	v[3] -= o[3]
}

func (v *V4) SubMutf(f Float) {
	v[0] -= f
	v[1] -= f
	v[2] -= f
	v[3] -= f
}

func (v V4) Mul(o V4) V4 {
	return V4{v[0] * o[0], v[1] * o[1], v[2] * o[2], v[3] * o[3]}
}

func (v V4) Mulf(f Float) V4 {
	return V4{v[0] * f, v[1] * f, v[2] * f, v[3] * f}
}

func (v *V4) MulMut(o V4) {
	v[0] *= o[0]
	v[1] *= o[1]
	v[2] *= o[2]
	v[3] *= o[3]
}

func (v *V4) MulMutf(f Float) {
	v[0] *= f
	v[1] *= f
	v[2] *= f
	v[3] *= f
}

func (v V4) Div(o V4) V4 {
	return V4{v[0] / o[0], v[1] / o[1], v[2] / o[2], v[3] / o[3]}
}

func (v V4) Divf(f Float) V4 {
	return V4{v[0] / f, v[1] / f, v[2] / f, v[3] / f}
}

func (v *V4) DivMut(o V4) {
	v[0] /= o[0]
	v[1] /= o[1]
	v[2] /= o[2]
	v[3] /= o[3]
}

func (v *V4) DivMutf(f Float) {
	v[0] /= f
	v[1] /= f
	v[2] /= f
	v[3] /= f
}

func (v V4) Eq(o V4) bool {
	return Abs(v[0]-o[0]) < FloatMin && Abs(v[1]-o[1]) < FloatMin && Abs(v[2]-o[2]) < FloatMin && Abs(v[3]-o[3]) < FloatMin
}

func (v V4) CloseEnough(o V4) bool {
	return CloseEnough(v[0], o[0]) && CloseEnough(v[1], o[1]) && CloseEnough(v[2], o[2]) && CloseEnough(v[3], o[3])
}

func (v V4) Dot(o V4) Float {
	return v[0]*o[0] + v[1]*o[1] + v[2]*o[2] + v[3]*o[3]
}

func (v V4) Mag() Float {
	return Sqrt(v.Dot(v))
}

func (v V4) Distance(to V4) Float {
	return v.Sub(to).Mag()
}

func (v V4) Angle(to V4) Angle {
	return FromRad(Acos(v.Dot(to) / (v.Mag() * to.Mag())))
}

func (v V4) Abs() V4 {
	return V4{Abs(v[0]), Abs(v[1]), Abs(v[2]), Abs(v[3])}
}

func (v V4) Normalize() V4 {
	return v.Mulf(1.0 / v.Mag())
}

func (v *V4) NormalizeMut() {
	v.MulMutf(1.0 / v.Mag())
}

func (v V4) Invert() V4 {
	return V4{-v[0], -v[1], -v[2], -v[3]}
}

func (v *V4) InvertMut() {
	v[0] = -v[0]
	v[1] = -v[1]
	v[2] = -v[2]
	v[3] = -v[3]
}

func (v V4) Clamp(min Float, max Float) V4 {
	return V4{Clamp(v[0], min, max), Clamp(v[1], min, max), Clamp(v[2], min, max), Clamp(v[3], min, max)}
}

func (v *V4) ClampMut(min Float, max Float) {
	v[0] = Clamp(v[0], min, max)
	v[1] = Clamp(v[1], min, max)
	v[2] = Clamp(v[2], min, max)
	v[3] = Clamp(v[3], min, max)
}

func (v V4) Lerp(to V4, t Float) V4 {
	return V4{Lerp(v[0], to[0], t), Lerp(v[1], to[1], t), Lerp(v[2], to[2], t), Lerp(v[3], to[3], t)}
}

func (v *V4) LerpMut(to V4, t Float) {
	v[0] = Lerp(v[0], to[0], t)
	v[1] = Lerp(v[1], to[1], t)
	v[2] = Lerp(v[2], to[2], t)
	v[3] = Lerp(v[3], to[3], t)
}

// Implements the [color.Color] interface
func (v V4) RGBA() (r, g, b, a uint32) {
	r = uint32(v[0] * 0xFFFF)
	g = uint32(v[1] * 0xFFFF)
	b = uint32(v[2] * 0xFFFF)
	a = uint32(v[3] * 0xFFFF)
	return
}

// Implements the [fmt.Stringer] interface
func (v V4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v[0], v[1], v[2], v[3])
}
