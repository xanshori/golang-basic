package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("continue")
			continue
		}else if i == 5{
			fmt.Println("break")
			break	
		}
		fmt.Println(i)
	}
}