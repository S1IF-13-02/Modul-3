package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	radius := float64(r)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	luas := 4 * math.Pi * math.Pow(radius, 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
