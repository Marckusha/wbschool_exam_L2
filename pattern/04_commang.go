package main

import "fmt"

type Command interface {
	Execute()
}

type ConcreteCommandA struct {
	rec *Receiver
}

func (c *ConcreteCommandA) Execute() {
	c.rec.Action("Command A")
}

type ConcreteCommandB struct {
	rec *Receiver
}

func (c *ConcreteCommandB) Execute() {
	c.rec.Action("Command B")
}

type Receiver struct{}

func (r *Receiver) Action(msg string) {
	fmt.Println(msg)
}

type Invoker struct {
	history []Command
}

func (inv *Invoker) StoreAndExecute(cmd Command) {
	inv.history = append(inv.history, cmd)
	cmd.Execute()
}

func main() {
	receiver := &Receiver{}
	cmdA := &ConcreteCommandA{receiver}
	cmdB := &ConcreteCommandB{receiver}
	inv := &Invoker{}
	inv.StoreAndExecute(cmdA)
	inv.StoreAndExecute(cmdB)

}
