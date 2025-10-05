package main

import (
	"fmt"
	"math"
)

func main() {
	var Jejari int
	fmt.Print("Jejari: ")
	fmt.Scanln(&Jejari)

	r := float64(Jejari)
	pi := math.Pi

	volume := (4.0 / 3.0) * pi * math.Pow(r, 3)
	luasKulit := 4 * pi * math.Pow(r, 2)

	fmt.Printf("Bola dengan Jejari %d memiliki volume %.4f dan luas kulit %.4f\n", Jejari, volume, luasKulit)

}
