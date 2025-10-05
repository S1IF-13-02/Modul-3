package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	fmt.Print("Masukkan jejari bola: ")
	fmt.Scan(&r)

	const pi = 3.1415926535
	radius := float64(r)
	volume := (4.0 / 3.0) * pi * math.Pow(radius, 3)
	luas := 4 * pi * math.Pow(radius, 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
