package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a int
	var mx sync.Mutex

	inc10000 := func() {
		for i := 0; i < 10000; i++ {
			mx.Lock()
			a++
			mx.Unlock()
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
