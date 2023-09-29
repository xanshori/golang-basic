package main

import (
	"fmt"
)

func main() {
	// Input
	var input string
	fmt.Print("Masukkan sesuatu: ")
	fmt.Scanln(&input)

	// Output
	fmt.Printf("Anda memasukkan: %s\n", input)
	fmt.Scanln()

}
