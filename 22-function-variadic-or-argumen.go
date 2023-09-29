package main

import "fmt"

func sumAll(number ...int) (total int) {
	for i := 0; i < len(number); i++ {
		total += number[i]
	}
	return
}

func main() {
	fmt.Println(sumAll(5,5,5,5))
	// or
	numbers :=[]int{5,5,5,5,20,}
	fmt.Println(sumAll(numbers...))
}