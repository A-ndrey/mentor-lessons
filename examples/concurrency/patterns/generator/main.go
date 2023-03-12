package main

import "fmt"

func main() {
	for v := range Range(5, 10) {
		fmt.Println(v)
	}
}

func Range(start, end int) <-chan int {
	c := make(chan int)
	go func() {
		for val := start; val < end; val++ {
			c <- val
		}
		close(c)
	}()

	return c
}
