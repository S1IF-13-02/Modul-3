package main

import "fmt"

func main() {
	var fx float64

	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	x := (2 / (fx - 5)) - 5

	fmt.Printf("Nilai x adalah: %.3f\n", x)
}
