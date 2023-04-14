package funcs

import "sync"

const size = 1000

func BigSlice() []int64 {
	res := make([]int64, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, int64(i))
	}

	return res
}

func SmallSlice() []int8 {
	res := make([]int8, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, int8(i%256))
	}

	return res
}

var space []int64
var spaceLock sync.Mutex

func Leak() {
	spaceLock.Lock()
	defer spaceLock.Unlock()
	for i := 0; i < size; i++ {
		space = append(space, int64(i))
	}
}
