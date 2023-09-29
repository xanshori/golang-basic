package main

import "fmt"

func main() {
	var hobi [3]string
	hobi[0] = "basket"
	hobi[1] = "bola"
	hobi[2] = "badminton"

	fmt.Println(hobi[0])
	fmt.Println(hobi[1])
	fmt.Println(hobi[2])

	// atau bisa langsung membuatnya dengan banyak

	jurusan := [10] string{
		"Teknik INformatik",
		"Sistem Informasi",
		"Management Informasi",
	}
	jurusan[0] = "Teknik Informatika"
	fmt.Println(jurusan)
	// len jurusan menunjukan jumlah array pada variable bukan isi
	// dari contoh diatas maka panjang data adalah 10 walaupun isi dari aray cuman 3
	fmt.Println(len(jurusan))
	fmt.Println(jurusan[0])
	
	var kelas [] interface{} = [] interface{}{
		"ruangan 1",
		"ruangan 2",
		"ruangan 3",
	}
	fmt.Println(len(kelas))
	fmt.Println(kelas[0])
}