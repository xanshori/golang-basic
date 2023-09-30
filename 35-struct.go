package main

import "fmt"

func main() {
	// biasanya aturan penggunaan nama sama persis dengan python
	// misalnya NamaMahasiswa
	type Mahasiswa struct {
		Name, Alamat, Jurusan string
		Nim                   int
	}
	// bisa gunakan ini atau
	var fuat Mahasiswa
	fuat.Name = "Fuat Anshori"
	fuat.Alamat = "Banjarmasin"
	fuat.Jurusan = "Teknik Informatika"
	fuat.Nim = 20041066
	fmt.Println(fuat)

	// atau ini
	hayabusa := Mahasiswa{
		Name:    "Hayabusa",
		Alamat:  "Banjarbaru",
		Jurusan: "Teknik Informatika",
		Nim:     20041066,
	}
	fmt.Println(hayabusa)

	// jika anda ingin membuat data dan dimasukan kedalam array
	data_mahasiswa := []Mahasiswa{
		{Name: "Alice", Alamat: "Samarinda", Jurusan: "Teknik Informatika", Nim: 20041010},
		{Name: "John", Alamat: "Tanah Laut", Jurusan: "Teknik Elektro", Nim: 20041020},
		{Name: "Doe", Alamat: "Tanah Bumbu", Jurusan: "Teknik Mesin", Nim: 20041030},
	}
	for _, data := range data_mahasiswa {
		fmt.Println(data.Name)
	}
	fmt.Println(data_mahasiswa)
}
