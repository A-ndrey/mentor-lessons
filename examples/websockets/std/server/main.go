package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Echo(ws *websocket.Conn) {
	withLog := io.MultiWriter(os.Stdout, ws)

	go func() {
		var count int
		for {
			time.Sleep(time.Second)
			ws.Write([]byte(strconv.Itoa(count)))
			count++
		}
	}()
	_, _ = io.Copy(withLog, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(Echo))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	err = http.Serve(&MyListener{listener: ln}, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

type MyListener struct {
	listener net.Listener
}

func (m *MyListener) Accept() (net.Conn, error) {
	conn, err := m.listener.Accept()

	return &MyConn{conn: conn}, err
}

func (m *MyListener) Close() error {
	return m.listener.Close()
}

func (m *MyListener) Addr() net.Addr {
	return m.listener.Addr()
}

type MyConn struct {
	conn net.Conn
}

func (m *MyConn) Read(b []byte) (n int, err error) {
	n, err = m.conn.Read(b)
	os.Stdout.Write(b[:n])

	return n, err
}

func (m *MyConn) Write(b []byte) (n int, err error) {
	n, err = m.conn.Write(b)
	os.Stdout.Write(b[:n])

	return n, err
}

func (m *MyConn) Close() error {
	return m.conn.Close()
}

func (m *MyConn) LocalAddr() net.Addr {
	return m.conn.LocalAddr()
}

func (m *MyConn) RemoteAddr() net.Addr {
	return m.conn.RemoteAddr()
}

func (m *MyConn) SetDeadline(t time.Time) error {
	return m.conn.SetDeadline(t)
}

func (m *MyConn) SetReadDeadline(t time.Time) error {
	return m.conn.SetReadDeadline(t)
}

func (m *MyConn) SetWriteDeadline(t time.Time) error {
	return m.conn.SetWriteDeadline(t)
}
