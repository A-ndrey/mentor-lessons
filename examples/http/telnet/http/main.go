package main

import (
	"io"
	"log"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.Header)
	w.Header().Add("X-Server", "example-server")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Some text for body\n"))
	if err != nil {
		log.Println(err)
		return
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.Header)
	w.WriteHeader(http.StatusAccepted)
	_, err := io.Copy(w, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
}

func main() {
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/echo", echoHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
