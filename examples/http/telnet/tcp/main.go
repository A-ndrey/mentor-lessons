package main

import (
	"io"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	w := io.MultiWriter(conn, os.Stdout)
	_, err := io.Copy(w, conn)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}
}
