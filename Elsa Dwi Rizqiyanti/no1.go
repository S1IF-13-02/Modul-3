package main

import (
	"fmt"
	"math"
)

func main() {
	var fx float64

	// Input nilai f(x)
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Rumus: x = 2/(f(x) - 5) - 5
	x := 2/(fx-5) - 5

	// Dibulatkan ke integer terdekat
	fmt.Println(int(math.Round(x)))
}
