package main

import "fmt"

func cekTipe(data interface{}) {
	switch v := data.(type) {
	case string:
		fmt.Println("Ini string:", v)
	case int:
		fmt.Println("Ini integer:", v)
	case float64:
		fmt.Println("Ini float:", v)
	default:
		fmt.Println("Tipe lain:", v)
	}
}

func main() {
	cekTipe("Halo")
	cekTipe(100)
	cekTipe(3.14)
	cekTipe(true)
}
