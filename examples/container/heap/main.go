package main

import (
	"container/heap"
	"fmt"
)

type PQ []int

func (p PQ) Len() int {
	return len(p)
}

func (p PQ) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQ) Push(x any) {
	i := x.(int)
	*p = append(*p, i)
}

func (p *PQ) Pop() any {
	i := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]

	return i
}

func main() {
	var pq PQ

	pq = append(pq, 100, 200, 150, 120, 140)

	heap.Init(&pq)

	for pq.Len() > 0 {
		v := heap.Pop(&pq)
		fmt.Println(v)
	}

	heap.Push(&pq, 1)
	heap.Push(&pq, 2)
	heap.Push(&pq, 3)
	heap.Push(&pq, 10)
	heap.Push(&pq, 12)
	heap.Push(&pq, 13)
	heap.Push(&pq, 8)
	heap.Push(&pq, 7)
	heap.Push(&pq, 6)

	for pq.Len() > 0 {
		v := heap.Pop(&pq)
		fmt.Println(v)
	}

	for _, v := range []int{1, 100, 20, 30, 452, 234, 523} {
		heap.Push(&pq, v)
		if pq.Len() > 3+1 {
			heap.Pop(&pq)
		}
	}

	// n = 10
	// k = 4
	// for(heap.Push+heap.Pop) 		O(nlogk) 2*2*10 = 40
	// sort.Sort 			O(nlogn) 10*3 = 30

	// lru cache
}
