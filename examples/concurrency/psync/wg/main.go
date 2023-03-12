package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printMsg(&wg, "first")

	wg.Add(2)
	go printMsg(&wg, "second")
	go printMsg(&wg, "third")

	wg.Wait()

	fmt.Println("last")
}

func printMsg(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	fmt.Println(msg)
}
