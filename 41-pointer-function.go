package main

import "fmt"

type Human struct {
	Rambut, Kulit string
	Is_life       interface{}
	Tinggi, Berat int
}

func changeKulit(h *Human) {
	h.Kulit = "Putih"
}
func modifyValue(ptr *int) {
	*ptr = 100
}
func main() {

	cio := Human{Rambut: "Hitam", Kulit: "", Is_life: true, Tinggi: 170, Berat: 60}
	var humanpointer *Human = &cio
	changeKulit(humanpointer)
	fmt.Println(cio)

	x := 42
	fmt.Println("Sebelum fungsi:", x) // Output: Sebelum fungsi: 42
	modifyValue(&x)                   // Mengubah nilai variabel x melalui pointer
	fmt.Println("Setelah fungsi:", x) // Output: Setelah fungsi: 100

	a := 50
	var b *int = &a
	var c *int = &a
	*b = 10
	*c = 100
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)
}
