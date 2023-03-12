package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var a int32

	inc10000 := func() {
		for i := 0; i < 10000; i++ {
			atomic.AddInt32(&a, 1)
		}
	}

	go inc10000()
	go inc10000()
	go inc10000()
	go inc10000()
	go inc10000()

	time.Sleep(2 * time.Second)

	fmt.Println(a)
}
