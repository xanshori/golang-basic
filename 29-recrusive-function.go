package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	fmt.Println(1*2*3*4*5)
	fmt.Println(factorial(15))
	fmt.Scanln()
	
}