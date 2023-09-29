package main

import "fmt"

func penjumlahan(a int, b int) int {
	return a + b
}

func main() {
	var a int = 10
	var b int = 20
	var penjumlahan int = penjumlahan(a, b)
	fmt.Println(penjumlahan)
}