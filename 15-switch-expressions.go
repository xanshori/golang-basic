package main

import "fmt"

func main() {
	nama := "admina"
	switch nama {
	case "adi":
		fmt.Println("adi")
	case "jamal":
		fmt.Println("jamal")
	default:
		fmt.Println("gapunya nama")
	}

}