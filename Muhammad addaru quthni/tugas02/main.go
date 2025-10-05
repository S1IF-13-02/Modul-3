package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Printf("Masukkan jari-jari bola: ")
	fmt.Scan(&radius)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	luas := 4 * math.Pi * math.Pow(radius, 2)

	fmt.Printf("Volume bola: %.4f\n", volume)
	fmt.Printf("Luas permukaan bola: %.4f\n", luas)
}
