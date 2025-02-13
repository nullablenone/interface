package main

import "fmt"

type Makhluk interface {
	bernafas() string
}

type Berjalan interface {
	berjalan() string
}

type Berenang interface {
	berenang() string
}

type Manusia struct{}

func (m Manusia) bernafas() string {
	return "Manusia Bernafas!"
}
func (m Manusia) berjalan() string {
	return "Manusia Berjalan!"
}

type Hewan struct{}

func (h Hewan) bernafas() string {
	return "Hewan Bernafas!"
}
func (h Hewan) berenang() string {
	return "Hewan Berenang"
}

func isBernafas(m Makhluk) {
	fmt.Println(m.bernafas())
}

func isBerjalan(m Berjalan) {
	fmt.Println(m.berjalan())
}

func isBerenang(m Berenang) {
	fmt.Println(m.berenang())
}

func main() {
	var manusia Manusia
	var hewan Hewan

	isBernafas(manusia)
	isBerjalan(manusia)
	fmt.Println("")
	isBernafas(hewan)
	isBerenang(hewan)
}
