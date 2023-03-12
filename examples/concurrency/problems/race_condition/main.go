package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("go->")
		}()
		go func() {
			fmt.Println("main")
		}()
		time.Sleep(10 * time.Millisecond)
	}

}
