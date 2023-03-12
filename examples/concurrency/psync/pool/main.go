package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var a int32
	pool := sync.Pool{
		New: func() any {
			return atomic.AddInt32(&a, 1)
		},
	}

	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)
	go printFromPool(&pool)

	time.Sleep(2 * time.Second)

}

func printFromPool(pool *sync.Pool) {
	val := pool.Get()
	for val == nil {
		val = pool.Get()
	}
	defer pool.Put(val)
	time.Sleep(time.Microsecond)
	fmt.Println(val)
}
