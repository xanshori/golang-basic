package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// siapapun yang mengacu memori dibawah ini
	var address Address = Address{
		"Banjarmasin", "Kalimantan Selatan", "Indonesia",
	}
	address1 := &address
	address2  := &address
	// ketika menggunakan lambang * maka akan dipindah ke addres 1 sebagai pointer
	*address1 = Address{
		"Kapuas", "Kalimantan Tengah", "Indonesia",
	}


	fmt.Println(address)
	fmt.Println(address1)
	fmt.Println(address2)
}