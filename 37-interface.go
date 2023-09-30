package main

import (
	"fmt"
)

/*
interface adalah tipe data abstrak, dia tidak memiliki implementasi langsung
sebuah interface berisikan method
biasanya interface berupa kontrak
*/

type Calculator interface {
	penjumlahan() int
	perkalian() int
	
}
type Nilai struct {
	A, B int
}

func (n Nilai) penjumlahan() int {
	return n.A + n.B
}
func (n Nilai) perkalian() int {
	return n.A * n.B
}


type Shape interface{
	kelilingPersegi() int
}
type NilaiSisi struct {
	Sisi int
}
func (ns NilaiSisi) kelilingPersegi() int {
	return ns.Sisi * 4
}
func main() {

	var penjumlahan1 Calculator = Nilai{1,2}
	fmt.Println(penjumlahan1.penjumlahan())

	var k1 Shape = NilaiSisi{10}
	fmt.Println(k1.kelilingPersegi())

}
