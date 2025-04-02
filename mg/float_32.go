//go:build MG_USE_F32

package mg

import "math"

type Float = float32

const (
	FloatMax Float = math.MaxFloat32
	FloatMin Float = math.SmallestNonzeroFloat32
)

func FloatBits(x Float) uint32     { return math.Float32bits(x) }
func FloatFromBits(b uint32) Float { return math.Float32frombits(b) }

func Abs(x Float) Float    { return FloatFromBits(FloatBits(x) &^ (1 << 31)) }
func Max(x, y Float) Float { return Float(math.Max(float64(x), float64(y))) }
func Sqrt(x Float) Float   { return Float(math.Sqrt(float64(x))) }

func Sin(x Float) Float { return Float(math.Sin(float64(x))) }
func Cos(x Float) Float { return Float(math.Cos(float64(x))) }
func Tan(x Float) Float { return Float(math.Tan(float64(x))) }

func Asin(x Float) Float     { return Float(math.Asin(float64(x))) }
func Acos(x Float) Float     { return Float(math.Acos(float64(x))) }
func Atan(x Float) Float     { return Float(math.Atan(float64(x))) }
func Atan2(y, x Float) Float { return Float(math.Atan2(float64(y), float64(x))) }

func Mod(x, y Float) Float { return Float(math.Mod(float64(x), float64(y))) }
