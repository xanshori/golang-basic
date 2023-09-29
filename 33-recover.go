package main

import "fmt"

func endapp() {
	// ini adalah fungsinya recover ini menangkap pesan dari panic yang berada di fungsi runapp
	message := recover()
	if message != nil{
		fmt.Println("pesan -> ",message)
	}else{
		fmt.Println("aplikasi berjalan")
	}
}

func runapp(error bool){
	// tujuanya mengabaikan error dan akan dipanggil ketika
	// fungsi ini selesai dijalankan atau berada dipaling akhir
	defer endapp()
	if error{
		panic("Terjadi Kesalahan")
	}
	// tidak akan di eksekusi jika error adalah true
	// tetapi jika kita menggunakan recover pada fungsi -
	// end app maka setelah program panic maka akan tetap berjalan
	
}
func main() {
	runapp(true)
}