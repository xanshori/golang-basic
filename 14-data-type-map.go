package main

import "fmt"

func main() {

	// map ["tipekey"] "tipevalue"
	// interface menunjukan bahwa tidak ada 
	// tipe data khusus untuk mengaturnya

	datas := map[interface{}]interface{}{
		1: "Fuat Anshori",
	}
	fmt.Println(datas)

	data := map[string]string{
		"j": "Fuat Anshori",
	}
	fmt.Println(datas,data)
}