package main

import (
	"golang.org/x/net/websocket"
	"log"
	"strconv"
)

func main() {
	httpUrl := "http://localhost/"
	wsUrl := "ws://localhost:8080/echo"
	ws, err := websocket.Dial(wsUrl, "", httpUrl)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 10; i++ {
		ws.Write([]byte("msg1" + strconv.Itoa(i)))
	}

	ws.Close()

	//go func() {
	//	io.Copy(ws, os.Stdin)
	//}()
	//
	//io.Copy(os.Stdout, ws)
}
