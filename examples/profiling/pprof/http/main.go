package main

import (
	"log"
	"mentor-space/examples/profiling/pprof/funcs"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	defaultHandler()
}

func defaultHandler() {

	http.HandleFunc("/", func(_ http.ResponseWriter, _ *http.Request) {
		funcs.SortSlice(false)
		funcs.SortSlice(true)
		funcs.Leak()
	})

	log.Println("http://localhost:8080/debug/pprof")

	log.Println(http.ListenAndServe("localhost:8080", nil))
}
