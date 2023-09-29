package main

import "fmt"

//  f(x) = 7x + 6 
func f(x int)int{
	return 7 * x + 6
}

//  g(x) = 4x – 6
func g(x int)int{
	return 4 * x - 6
}

//  (f o g)(x) 
func f_komposisi_g(x int)(result int){

	result = f(g(x)) 
	result = 7 * g(x) + 6
	result = 7 * (4 * x - 6) + 6
	result = 28 * x - 42  + 6
	result = 28 * x - 36
	return
}

func main() {
	fmt.Println(f_komposisi_g(10))
}