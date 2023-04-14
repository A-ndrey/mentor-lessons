package main

import (
	"log"
	"os"
	"runtime/pprof"
)

const size = 1_000_000

func main() {
	f, err := os.Create("pprof.cpu.out")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	m := createMap()
	s := createSlice()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatalln(err)
	}

	var sum = 0

	for _, v := range m {
		sum += v
	}

	for _, v := range s {
		sum += v
	}

	pprof.StopCPUProfile()

	log.Println(sum)

}

func createMap() map[int]int {
	res := make(map[int]int, size)
	for i := 0; i < size; i++ {
		res[i] = i + 1
	}

	return res
}

func createSlice() []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = i + 1
	}

	return res
}
