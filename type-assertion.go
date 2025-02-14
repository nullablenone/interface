package main

import "fmt"

func cetakData(data interface{}) {
	// Kalau yakin sama tipe datanya, kita bisa langsung konversi pakai type assertion.

	// Type assertion ke string
	str, ok := data.(string)
	if ok {
		fmt.Println("Ini string:", str)
	} else {
		fmt.Println("Bukan string!")
	}
}

func main() {
	cetakData("Halo") 
	cetakData(100)    
}
