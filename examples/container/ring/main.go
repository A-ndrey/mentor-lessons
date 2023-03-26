package main

import (
	"container/ring"
	"fmt"
)

func main() {

	rr := ring.New(3)
	rr.Value = "s1"
	rr.Next().Value = "s2"
	rr.Next().Next().Value = "s3"

	for i := 0; i < 100; i++ {
		fmt.Println(rr.Value)
		rr = rr.Next()
	}

	rr.Do(func(a any) {
		fmt.Printf("access ok %v", a)
	})
}
