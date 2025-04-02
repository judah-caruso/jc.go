package thelp

import "testing"

func Expect(t *testing.T, cond bool) {
	t.Helper()
	Expectf(t, cond, "expectation failed")
}

func Expectf(t *testing.T, cond bool, format string, args ...any) {
	t.Helper()
	if !cond {
		t.Fatalf(format, args...)
	}
}
