package main

import "fmt"

func main() {
	// constant ini digunakan agar variable tidak dapat dirubah
	const nama = "Fuat Anshori"
	fmt.Println(nama)
	// kita bisa juga menggunakan multiple constant

	const (
		umur = 20
		kelas = "IT"
	)
	fmt.Println(umur)
	fmt.Println(kelas)
}