package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	const pi = 3.1415926535

	fmt.Print("Jarijari = ")
	fmt.Scanln(&r)

	volume := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)
	luas := 4 * pi * math.Pow(float64(r), 2)

	fmt.Printf("Bola dengan jarijari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
