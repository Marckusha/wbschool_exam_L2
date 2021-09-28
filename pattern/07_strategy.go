package main

import "fmt"

type Navigator struct {
	route func()
}

func (n *Navigator) BuildRoute() {
	n.route()
}

func (n *Navigator) SetRoute(r func()) {
	n.route = r
}

func main() {
	carRoute := func() {
		fmt.Println("Build cat route")
	}
	manRoute := func() {
		fmt.Println("Build man route")
	}

	navigator := &Navigator{}
	navigator.SetRoute(carRoute)
	navigator.BuildRoute()
	navigator.SetRoute(manRoute)
	navigator.BuildRoute()
}
