package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.1415926535

	var r int
	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	rFloat := float64(r)

	volume := (4.0 / 3.0) * pi * rFloat * rFloat * rFloat
	luas := 4 * pi * rFloat * rFloat

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
