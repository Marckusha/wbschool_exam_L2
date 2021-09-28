package main

import "fmt"

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct{}

func (el *ConcreteElementA) Accept(v Visitor) {
	fmt.Println("Concrete element A")
	v.VisitA()
}

type ConcreteElementB struct{}

func (el *ConcreteElementB) Accept(v Visitor) {
	fmt.Println("Concrete elemnt B")
	v.VisitB()
}

type Visitor interface {
	VisitA()
	VisitB()
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitA() {
	fmt.Println("Cocnrete visitor A")
}

func (v *ConcreteVisitor) VisitB() {
	fmt.Println("Cocnrete visitor B")
}

func main() {
	visitor := &ConcreteVisitor{}
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
