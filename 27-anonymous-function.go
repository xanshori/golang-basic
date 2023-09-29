package main
import "fmt"

type BlackList func(string) bool
func registerUSer(name string, filter BlackList){
	if filter(name){
		fmt.Println("your name is blocked")
	}else{
		fmt.Println("success register")
	}
}

// function
func filterUsername(name string)bool{
	if name == "jamal"{
		return true
	}
	return false
}


func main() {
	// this is an anonymous function
	registerUSer("admin",func(name string)bool{return name=="admin"})
}