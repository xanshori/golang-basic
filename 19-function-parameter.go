package main

import "fmt"

func biodata(nama string, nim int) {
	fmt.Println("Hallo, my name is ",nama,"nim saya adalah ",nim)
}

func main() {
	name := "Fuat Anshori"
 biodata(name,20041066)
}