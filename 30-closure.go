package main

import "fmt"

func main() {

	name := "Fuat Anshori"
	myfunc := func() {
		// disarankan untuk deklarasi variable baru 
		// misalnya  menggunakan name:= "jamal"
		name = "Andi"
		
		fmt.Println("hello")
	}
	myfunc()
	fmt.Println(name)

}