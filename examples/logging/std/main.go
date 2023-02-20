package main

import (
	"log"
	"os"
)

func main() {
	//    ex1()
	//    ex2(); log.Println("after end")
	//    ex3(); log.Println("after end")
	//    ex4()
	ex5()
}

func ex1() {
	log.Println("Hello, new line")
	log.Print("Hello")
	log.Print("log\n")

	data := struct {
		Str   string
		Int   int
		Slice []int
	}{
		Str:   "test string",
		Int:   200,
		Slice: []int{0, 100, 200, 300},
	}
	log.Printf("Hello, data: %v", data)
}

func ex2() {
	log.Fatalln("Something went wrong")
}

func ex3() {
	log.Panic("Something went wrong")
}

func ex4() {
	logger := log.New(os.Stdout, "new logger", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("new message")

	logger = log.New(os.Stdout, "new logger 2", log.Ldate|log.Lshortfile|log.Lmsgprefix)
	logger.Println("new message 2")

	logger.SetFlags(log.LstdFlags)
	logger.Println("new message 3")
}

func ex5() {
	log.SetFlags(log.Flags() | log.Llongfile)
	_ = log.Output(0, "print something")
}
