package main

import "fmt"

func endApp() {
	fmt.Println("app ended")
}

func runApp(error bool) {
	// tujuanya mengabaikan error dan akan dipanggil ketika
	// fungsi ini selesai dijalankan atau berada dipaling akhir
	defer endApp()
	if error{
		panic("aplikasi err")
	}
	// tidak akan di eksekusi jika error adalah true
	fmt.Println("Aplikasi Berjalan")
}
func main() {
	runApp(true)
}
