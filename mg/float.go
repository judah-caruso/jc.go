package mg

const Epsilon = 0.000001

func Clamp(x, minimum, maximum Float) Float {
	return min(maximum, max(minimum, x))
}

func InvSqrt(x Float) Float {
	return 1 / Sqrt(x)
}

func Lerp(a, b, t Float) Float {
	return (1.0-t)*a + t*b
}

func CloseEnough(a, b Float) bool {
	return Abs(a-b) <= (Epsilon * Max(1.0, Max(Abs(a), Abs(b))))
}
