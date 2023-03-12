package main

import (
	"fmt"
	"time"
)

func main() {
	var a int

	inc10000 := func() {
		for i := 0; i < 10000; i++ {
			a++
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
