package main

import (
	"fmt"
	"os"
)

func main() {
	args :=os.Args
	fmt.Println("argument")
	fmt.Println(args)
	fmt.Println(args[0])
	name,err := os.Hostname()
	if err == nil{
		fmt.Println("Hostname : "+name)
	}else{
		fmt.Println("Error : "+err.Error())

	}
	fmt.Println(os.Getenv("USERNAME"))
}