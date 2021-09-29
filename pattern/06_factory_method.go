package main

import "fmt"

type Button interface {
	render()
}

type OSXButton struct{}

func (osx *OSXButton) render() {
	fmt.Println("OSX button")
}

type WindowsButton struct{}

func (w *WindowsButton) render() {
	fmt.Println("Windows button")
}

type GUIFactory interface {
	createButton() Button
}

type OSXFactory struct{}

func (f *OSXFactory) createButton() Button {
	return &OSXButton{}
}

type WindowsFactory struct{}

func (f *WindowsFactory) createButton() Button {
	return &WindowsButton{}
}

type App struct {
	guiFact GUIFactory
}

func (c *App) Operation() {
	guiElement := c.guiFact.createButton()
	guiElement.render()
}

func main() {
	application := &App{&WindowsFactory{}}
	application.Operation()
}
