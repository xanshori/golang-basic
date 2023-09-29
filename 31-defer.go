package main

import "fmt"

func logging() {
	fmt.Println("fucntion called")
}

func runApplication(value int){
	// tujuanya mengabaikan error dan akan dipanggil ketika
	// fungsi ini selesai dijalankan atau berada dipaling akhir
	defer logging()
	result := 20/value
	fmt.Println(result)
	
}
func main() {
	runApplication(0)
}