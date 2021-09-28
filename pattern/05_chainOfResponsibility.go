package main

import "fmt"

type Handler interface {
	Request(fl bool)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) Request(fl bool) {
	fmt.Println("ConcreteHandler A")
	if fl {
		h.next.Request(fl)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) Request(fl bool) {
	fmt.Println("ConcreteHandler B")
}

func main() {
	handlA := &ConcreteHandlerA{new(ConcreteHandlerB)}
	handlA.Request(true)
}
