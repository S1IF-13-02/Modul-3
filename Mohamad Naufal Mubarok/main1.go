package main

import (
	"fmt"
)

func main() {
	var y float64
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&y)

	
	x := (2.0 / (y - 5)) - 5

	fmt.Printf("Nilai x adalah: %.0f\n", x)
}