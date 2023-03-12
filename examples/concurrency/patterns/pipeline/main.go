package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	pipeline()
	fmt.Println(time.Since(now))

	//now := time.Now()
	//commonFlow()
	//fmt.Println(time.Since(now))
}

func pipeline() {
	gen := Generator()

	p1 := Processor(gen, process("p1"))
	p2 := Processor(p1, process("p2"))
	p3 := Processor(p2, process("p3"))
	p4 := Processor(p3, process("p4"))

	Print(p4)
	// 5     4      3     2    1      0
	//gen -> p1 -> p2 -> p3 -> p4 -> Print()
}

func commonFlow() {
	for v := range Generator() {
		v = process("p1")(v)
		v = process("p2")(v)
		v = process("p3")(v)
		v = process("p4")(v)
		fmt.Println(v)
	}
}

func process(suffix string) func(string) string {
	return func(s string) string {
		time.Sleep(100 * time.Millisecond)
		return s + " " + suffix
	}
}

func Generator() chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 20; i++ {
			c <- strconv.Itoa(i)
		}

		close(c)
	}()

	return c
}

func Processor(val chan string, f func(string) string) chan string {
	c := make(chan string)
	go func() {
		for v := range val {
			c <- f(v)
		}

		close(c)
	}()

	return c
}

func Print(val chan string) {
	for v := range val {
		fmt.Println(v)
	}
}
