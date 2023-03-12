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
	time.Sleep(time.Second)
}

func lockPrint(mx1, mx2 *sync.Mutex, msg string) {
	mx1.Lock()
	defer mx1.Unlock()
	time.Sleep(time.Millisecond)
	mx2.Lock()
	defer mx2.Unlock()
	fmt.Println(msg)
}
