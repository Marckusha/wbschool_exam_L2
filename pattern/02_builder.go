package main

import "strconv"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type NonPayerCharacter struct {
	name      string
	charClass string
	loots     []string
}

func (npc *NonPayerCharacter) String() string {
	return "Name character: " + npc.name + ", count loots: " + strconv.Itoa(len(npc.loots)) + ", character class: " + npc.charClass
}

type Director struct {
	builder Builder
}

func NewDirector(b *Builder) *Director {
	return &Director{
		builder: *b,
	}
}

func (d *Director) Construct() *NonPayerCharacter {
	d.builder.SetName()
	d.builder.SetCharClass()
	d.builder.SetLoots()
	return d.builder.GetResult()
}

type Builder interface {
	SetName()
	SetCharClass()
	SetLoots()
	GetResult() *NonPayerCharacter
}

type MageBuilder struct {
	npc *NonPayerCharacter
}

func NewMageBuilder() *MageBuilder {
	return &MageBuilder{new(NonPayerCharacter)}
}

func (m *MageBuilder) SetName() {
	m.npc.name = "armagedone"
}

func (m *MageBuilder) SetCharClass() {
	m.npc.charClass = "mage warrior"
}

func (m *MageBuilder) SetLoots() {
	m.npc.loots = append(m.npc.loots, "staff", "wizard hat")
}

func (m *MageBuilder) GetResult() *NonPayerCharacter {
	return m.npc
}
