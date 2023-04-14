package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		_, err := http.Get("http://localhost:8080/")
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
}
