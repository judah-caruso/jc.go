//go:build !MG_USE_F32

package mg

import "math"

type Float = float64

const (
	FloatMax Float = math.MaxFloat64
	FloatMin Float = math.SmallestNonzeroFloat64
)

func FloatBits(x Float) uint64     { return math.Float64bits(x) }
func FloatFromBits(b uint64) Float { return math.Float64frombits(b) }

func Abs(x Float) Float    { return math.Abs(x) }
func Max(x, y Float) Float { return math.Max(x, y) }
func Sqrt(x Float) Float   { return math.Sqrt(x) }

func Sin(x Float) Float { return math.Sin(x) }
func Cos(x Float) Float { return math.Cos(x) }
func Tan(x Float) Float { return math.Tan(x) }

func Asin(x Float) Float     { return math.Asin(x) }
func Acos(x Float) Float     { return math.Acos(x) }
func Atan(x Float) Float     { return math.Atan(x) }
func Atan2(y, x Float) Float { return math.Atan2(y, x) }

func Mod(x, y Float) Float { return math.Mod(x, y) }
