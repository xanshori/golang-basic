// helper.go
package helper

import "fmt"

func SayHello(name string) {
    fmt.Println("Haii " + name)
}
// menggunakan huruf kecil pada awal nama function
// menandakan function tidak di export diluar package
func getCar(name string) {
    fmt.Println("Haii " + name)
}
