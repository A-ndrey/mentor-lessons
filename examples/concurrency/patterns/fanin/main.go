package main

import (
	"fmt"
	"sync"
)

func main() {
	c := FanIn(
		Generator(1, 5),
		Generator(6, 10),
		Generator(100, 110),
	)

	for v := range c {
		fmt.Println(v)
	}

	//                                  -> Processor              ->
	// Generator -> Processor -> FanOut -> Processor              -> FanIn -> Processor
	//                                  -> Processor -> Processor ->
}

func Generator(start, end int) chan int {
	c := make(chan int)
	go func() {
		for i := start; i < end; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func FanIn(cs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	send := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
