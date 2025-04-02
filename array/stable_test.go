package array_test

import (
	"runtime"
	"testing"

	"github.com/judah-caruso/jc.go/array"
	"github.com/judah-caruso/jc.go/thelp"
)

func TestStableWithGC(t *testing.T) {
	type valuewithptr struct {
		value int
		ptr   *int
	}

	var arr array.Stable[valuewithptr]
	arr.Init()

	aptr := arr.Append(valuewithptr{value: 10, ptr: nil})
	bptr := arr.Append(valuewithptr{value: 20, ptr: &aptr.value})

	for i := range 100 {
		arr.Append(valuewithptr{value: i})
		runtime.GC()
	}

	thelp.Expect(t, arr.Get(0) == aptr)
	thelp.Expect(t, arr.Get(1) == bptr)
	thelp.Expectf(t, arr.Len() == 102, "len was %d", arr.Len())
	thelp.Expect(t, bptr.ptr != nil && bptr.value == 20)
	thelp.Expectf(t, bptr.ptr == &aptr.value, "%p vs. %p", bptr.ptr, &aptr.value)
}
