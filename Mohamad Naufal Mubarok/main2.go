package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	pi := 3.1415926535
	volume := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)
	luas := 4 * pi * math.Pow(float64(r), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
