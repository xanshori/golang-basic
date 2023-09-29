package main

import "fmt"

func main() {
	a := 10
	b := 190
	if a == b {
		fmt.Println("if 1")
	}else if a>b{
		fmt.Println("if 2")
	}else{
		fmt.Println("else")
	}
}