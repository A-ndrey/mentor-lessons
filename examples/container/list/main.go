package main

import (
	"container/list"
	"fmt"
)

func main() {
	q := list.New()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	for q.Len() > 0 {
		v := q.Front()
		fmt.Println(v.Value)
		q.Remove(v)
		if v.Value.(int)%2 == 0 {
			q.PushBack(v.Value.(int)*100 - 1)
		}
	}
}
