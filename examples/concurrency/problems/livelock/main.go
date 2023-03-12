package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mx1, mx2 sync.Mutex
	go lockPrint(&mx1, &mx2, "first")
	go lockPrint(&mx2, &mx1, "second")
	time.Sleep(5 * time.Second)
}

func lockPrint(mx1, mx2 *sync.Mutex, msg string) {
	for !mx1.TryLock() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(msg + " waiting for mx1")
	}
	defer mx1.Unlock()
	time.Sleep(time.Millisecond)
	for !mx2.TryLock() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(msg + " waiting for mx2")
	}
	defer mx2.Unlock()
	fmt.Println(msg)
}
