package main

import "fmt"

func nameFilter(name string, wordFilter func(string) string) {
	fmt.Println(wordFilter(name))
}

func wordFilter(name string) string{
	if name == "fuck"{
		return "f*ck"
	}
	return name
}

func main() {
	nameFilter("fuck",wordFilter)
}