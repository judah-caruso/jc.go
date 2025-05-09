package mg

import "fmt"

// V2 is a 2-dimensional vector type
type V2 [2]Float

func (v V2) X() Float      { return v[0] }
func (v V2) Y() Float      { return v[1] }
func (v V2) U() Float      { return v[0] }
func (v V2) V() Float      { return v[1] }
func (v V2) Min() Float    { return v[0] }
func (v V2) Max() Float    { return v[1] }
func (v V2) Width() Float  { return v[0] }
func (v V2) Height() Float { return v[1] }

func (v *V2) Xp() *Float      { return &v[0] }
func (v *V2) Yp() *Float      { return &v[1] }
func (v *V2) Up() *Float      { return &v[0] }
func (v *V2) Vp() *Float      { return &v[1] }
func (v *V2) Minp() *Float    { return &v[0] }
func (v *V2) Maxp() *Float    { return &v[1] }
func (v *V2) Widthp() *Float  { return &v[0] }
func (v *V2) Heightp() *Float { return &v[1] }

func (v *V2) SetX(f Float)      { v[0] = f }
func (v *V2) SetY(f Float)      { v[1] = f }
func (v *V2) SetU(f Float)      { v[0] = f }
func (v *V2) SetV(f Float)      { v[1] = f }
func (v *V2) SetMin(f Float)    { v[0] = f }
func (v *V2) SetMax(f Float)    { v[1] = f }
func (v *V2) SetWidth(f Float)  { v[0] = f }
func (v *V2) SetHeight(f Float) { v[1] = f }

func (v V2) Add(o V2) V2 {
	return V2{v[0] + o[0], v[1] + o[1]}
}

func (v V2) Addf(f Float) V2 {
	return V2{v[0] + f, v[1] + f}
}

func (v *V2) AddMut(o V2) {
	v[0] += o[0]
	v[1] += o[1]
}

func (v *V2) AddMutf(f Float) {
	v[0] += f
	v[1] += f
}

func (v V2) Sub(o V2) V2 {
	return V2{v[0] - o[0], v[1] - o[1]}
}

func (v V2) Subf(f Float) V2 {
	return V2{v[0] - f, v[1] - f}
}

func (v *V2) SubMut(o V2) {
	v[0] -= o[0]
	v[1] -= o[1]
}

func (v *V2) SubMutf(f Float) {
	v[0] -= f
	v[1] -= f
}

func (v V2) Mul(o V2) V2 {
	return V2{v[0] * o[0], v[1] * o[1]}
}

func (v V2) Mulf(f Float) V2 {
	return V2{v[0] * f, v[1] * f}
}

func (v *V2) MulMut(o V2) {
	v[0] *= o[0]
	v[1] *= o[1]
}

func (v *V2) MulMutf(f Float) {
	v[0] *= f
	v[1] *= f
}

func (v V2) Div(o V2) V2 {
	return V2{v[0] / o[0], v[1] / o[1]}
}

func (v V2) Divf(f Float) V2 {
	return V2{v[0] / f, v[1] / f}
}

func (v *V2) DivMut(o V2) {
	v[0] /= o[0]
	v[1] /= o[1]
}

func (v *V2) DivMutf(f Float) {
	v[0] /= f
	v[1] /= f
}

func (v V2) Eq(o V2) bool {
	return Abs(v[0]-o[0]) < FloatMin && Abs(v[1]-o[1]) < FloatMin
}

func (v V2) CloseEnough(o V2) bool {
	return CloseEnough(v[0], o[0]) && CloseEnough(v[1], o[1])
}

func (v V2) Dot(o V2) Float {
	return v[0]*o[0] + v[1]*o[1]
}

func (v V2) Mag() Float {
	return Sqrt(v.Dot(v))
}

func (v V2) Distance(to V2) Float {
	return v.Sub(to).Mag()
}

func (v V2) Angle(to V2) Angle {
	return FromRad(Acos(v.Dot(to) / (v.Mag() * to.Mag())))
}

func (v V2) Abs() V2 {
	return V2{Abs(v[0]), Abs(v[1])}
}

func (v V2) Normalize() V2 {
	return v.Mulf(1.0 / v.Mag())
}

func (v *V2) NormalizeMut() {
	v.MulMutf(1.0 / v.Mag())
}

func (v V2) Invert() V2 {
	return V2{-v[0], -v[1]}
}

func (v *V2) InvertMut() {
	v[0] = -v[0]
	v[1] = -v[1]
}

func (v V2) Clamp(min Float, max Float) V2 {
	return V2{Clamp(v[0], min, max), Clamp(v[1], min, max)}
}

func (v *V2) ClampMut(min Float, max Float) {
	v[0] = Clamp(v[0], min, max)
	v[1] = Clamp(v[1], min, max)
}

func (v V2) Lerp(to V2, t Float) V2 {
	return V2{Lerp(v[0], to[0], t), Lerp(v[1], to[1], t)}
}

func (v *V2) LerpMut(to V2, t Float) {
	v[0] = Lerp(v[0], to[0], t)
	v[1] = Lerp(v[1], to[1], t)
}

func (v V2) Rotate(angle Angle) V2 {
	cos := Float(angle.Cos())
	sin := Float(angle.Sin())
	return V2{
		v[0]*cos - v[1]*sin,
		v[0]*sin + v[1]*cos,
	}
}

func (v *V2) RotateMut(angle Angle) {
	cos := Float(angle.Cos())
	sin := Float(angle.Sin())
	v[0] = v[0]*cos - v[1]*sin
	v[1] = v[0]*sin + v[1]*cos
}

func (v V2) Reflect(normal V2) V2 {
	dot := v.Dot(normal)
	return V2{
		v[0] - (2*normal[0])*dot,
		v[1] - (2*normal[1])*dot,
	}
}

func (v *V2) ReflectMut(normal V2) {
	dot := v.Dot(normal)
	v[0] = v[0] - (2*normal[0])*dot
	v[1] = v[1] - (2*normal[1])*dot
}

// Implements the [fmt.Stringer] interface
func (v V2) String() string {
	return fmt.Sprintf("(%f, %f)", v[0], v[1])
}
