package main

import (
	"fmt"
)

func random() interface{} {
	return 20
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("integer", value)
	}

}
