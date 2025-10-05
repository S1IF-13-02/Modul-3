package main

import (
	"fmt"
)

func main() {
	var fx float64

	// Input nilai f(x)
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Hitung nilai x berdasarkan persamaan f(x) = 2/(x-5) - 5
	x := 5 - 2/(fx-5)

	// Output nilai x sebagai integer
	fmt.Printf("Nilai x adalah: %d\n", int(x))
}