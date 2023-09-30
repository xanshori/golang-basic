package main

import "fmt"

type Data struct {
	Nama, Alamat string
	No           int
}


func (data Data) data2(){
	fmt.Println("hallo my name is ",data.Nama,)

}
func main() {
	nama2 := "fuat Anshori"
	nama_data2 := Data{Nama: nama2,Alamat: "Banjarbaru",No: 12}
	nama_data2.data2()
	fmt.Println(nama_data2.Nama)
	fmt.Println(nama_data2.Alamat)
	fmt.Println(nama_data2.No,"\n")

	nama3 := "Jamal"
	nama_data3 := Data{Nama: nama3,Alamat: "Banjarmasin",No: 12}
	nama_data3.data2()
	fmt.Println(nama_data3.Nama)
	fmt.Println(nama_data3.Alamat)
	fmt.Println(nama_data3.No)
}