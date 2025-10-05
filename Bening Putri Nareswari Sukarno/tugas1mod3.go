package main

import (
	"fmt"
)

func main() {
	var z float64
	fmt.Print("masukkan nilai f(x): ")
	fmt.Scan(&z)

	x := (2.0 / (z - 5)) - 5

	fmt.Printf("Nilai x adalah: %.2f\n", x)
}
