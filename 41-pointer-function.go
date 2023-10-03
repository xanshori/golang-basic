package main

import "fmt"

type Human struct {
	Rambut, Kulit string
	Is_life       interface{}
	Tinggi, Berat int
}

func changeTinggi(h *Human) {
	h.Kulit = "Putih"
}

func main() {
	cio := Human{"Hitam", "", true, 170, 60}
	changeTinggi(&cio)
	fmt.Println(cio)

}
 