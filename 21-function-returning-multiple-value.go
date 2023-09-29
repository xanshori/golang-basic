package main

import "fmt"

func pembagian(a int,b int)(result int) {
	result = a/b
	return 
}

func main() {
	fmt.Println(pembagian(4,2))
}