package main
import "fmt"


type Filter func(string) string
func nameFilter(name string, wordFilter Filter) {
	fmt.Println(wordFilter(name))
}


func wordFilter(name string) string{
	if name == "fuck"{
		return "f*ck"
	}else if name == "bitch" {
		return "b*tch"
	}
	return name
}


func main() {
	nameFilter("bitch",wordFilter)
}