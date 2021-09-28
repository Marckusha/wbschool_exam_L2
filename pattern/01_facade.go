package main

import (
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type CPU struct{}

func (c *CPU) Execute() {
	fmt.Println("CPU Execute")
}

type Memory struct{}

func (m *Memory) Load(data []byte) {
	fmt.Println("Load mem")
}

type HardDrive struct{}

func (hd *HardDrive) Read() []byte {
	fmt.Println("HardDrive read")
	return make([]byte, 0)
}

type ComputerFacade struct {
	proccesor *CPU
	ram       *Memory
	hd        *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		proccesor: new(CPU),
		ram:       new(Memory),
		hd:        new(HardDrive),
	}
}

func (c *ComputerFacade) Start() {
	c.proccesor.Execute()
	c.ram.Load(c.hd.Read())
}

func main() {
	comp := NewComputerFacade()
	comp.Start()
}
