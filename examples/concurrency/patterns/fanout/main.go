package main

import (
	"fmt"
	"time"
)

func main() {
	cs := FanOut(Generator(), 2)

	go func() {
		for v := range cs[0] {
			fmt.Println(v)
		}
	}()

	go func() {
		for v := range cs[1] {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second)
}

func Generator() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func FanOut(ch chan int, n int) []chan int {
	cs := make([]chan int, 0, n)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	toChannels := func(ch chan int, cs []chan int) {
		defer func(cs []chan int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}

	go toChannels(ch, cs)

	return cs
}
