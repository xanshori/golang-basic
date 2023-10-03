package main

import "fmt"

// ini adalah klas
type Man struct{
	Name string
}
// ini adalah method
func (man *Man) Merried(){
	man.Name = "Mr. "+man.Name
}

func main() {
	jamal := Man{"jamal"}
	jamal.Merried()
	fmt.Println(jamal.Name)
}