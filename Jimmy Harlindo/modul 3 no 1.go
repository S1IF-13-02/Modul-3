package main

import (
	"fmt"
)

func main() {
	var fx, x float64

	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	x = (2 / (fx - 5)) - 5

	fmt.Printf("Nilai x: %.0f\n", x)
}



