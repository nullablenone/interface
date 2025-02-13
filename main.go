package main

import "fmt"

type Makhluk interface {
	bernafas() string
}

type Manusia struct{}

func (m Manusia) bernafas() string {
	return "Manusia Bernafas!"
}

type Hewan struct{}

func (h Hewan) bernafas() string {
	return "Hewan Bernafas!"
}

func isBernafas(m Makhluk) {
	fmt.Println(m.bernafas())
}

func main() {
	var manusia Manusia
	var hewan Hewan

	isBernafas(manusia)
	isBernafas(hewan)
}
