package main

import "fmt"

func main() {
	names := []string{
		"jamal","asep","una","aliyah","toil",
	}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("nama : ", names[i])
	// }

	for idx, name := range names {
		fmt.Println("nama : ", name)
		fmt.Println("index ke : ", idx)
	}
}