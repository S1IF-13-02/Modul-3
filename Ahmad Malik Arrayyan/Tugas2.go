package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.1415926535

	var jari int
	fmt.Print("jari = ")
	fmt.Scan(&jari)

	volume := (4.0 / 3.0) * pi * math.Pow(float64(jari), 3)
	luas := 4 * pi * math.Pow(float64(jari), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", jari, volume, luas)
}