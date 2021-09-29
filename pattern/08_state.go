package main

import "fmt"

type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle()
}

func (c *Context) SetState(st State) {
	c.state = st
}

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (st *ConcreteStateA) Handle() {
	fmt.Println("This is state A")
}

type ConcreteStateB struct{}

func (st *ConcreteStateB) Handle() {
	fmt.Println("This is state B")
}

func main() {
	ctx := &Context{}
	stA := &ConcreteStateA{}
	stB := &ConcreteStateB{}
	ctx.SetState(stA)
	ctx.Request()
	ctx.SetState(stB)
	ctx.Request()
}
