package main

import "fmt"

func isType(tipe interface{}) {
	fmt.Printf("Type is : %T \n", tipe)
}
func main() {
	isType("muhamad yesa")
	isType(17)
	isType(true)

	// slice pake interface
	var dataCampuran []interface{} = []interface{}{"yesa", 17, 2007}
	fmt.Println(dataCampuran)

	cobaDiMap := map[string]interface{}{
		"nama":  "muhamad yesa",
		"umur":  17,
		"tahun": 2007,
	}
	fmt.Println(cobaDiMap)
}
