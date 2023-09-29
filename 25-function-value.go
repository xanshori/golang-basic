package main

import "fmt"

func nama(nama string) string {
	return "Hallo " + nama
}

func main() {
	nama_value_function := nama
	fmt.Println(nama_value_function("Fuat Anshori"))
}