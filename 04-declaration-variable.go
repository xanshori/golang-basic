package main

import "fmt"

func main() {
	var name string
	var umur int8

	name = "Fuat Anshori"
	fmt.Println(name)
	
	umur = 16
	fmt.Println(umur)

	var kelas = 12
	fmt.Println(kelas)

	// ketika tidak ingin menggunakan keyword "var" bisa diganti dengan ":="
	jurusan := "Ti"
	jurusan="Si"
	fmt.Println(jurusan)

	// kita juga bisa membuat multiple variable
	var (
		matakuliah = "IoT"
		sks = 1
	)
	fmt.Println(matakuliah)
	fmt.Println(sks)
}