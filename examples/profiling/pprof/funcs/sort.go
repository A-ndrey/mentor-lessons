package funcs

import (
	"log"
	"math/rand"
	"sort"
)

func SortSlice(fast bool) {
	data := randSlice(10000)

	if fast {
		fastSort(data)
	} else {
		slowSort(data)
	}

	checkCorrect(data)
}

func randSlice(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, rand.Int())
	}

	return res
}

func slowSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func fastSort(data []int) {
	sort.Ints(data)
}

func checkCorrect(data []int) {
	last := data[0]
	for _, v := range data[1:] {
		if last > v {
			log.Panicf("incorrect order %d gt %d", last, v)
		}
		last = v
	}
}
