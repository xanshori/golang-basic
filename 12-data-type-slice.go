package main

import "fmt"

func main() {
	// titik 3 dibawah ketika kita tidak mengetahui kapasitas
	var months = [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
	}
	slice1 := months[6:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
}