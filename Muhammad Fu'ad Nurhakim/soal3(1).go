package main

import "fmt"

func main() {
	var fx float64

	fmt.Scan(&fx)

	x := 2/(fx-5) - 5

	fmt.Printf("Nilai x adalah: %.0f\n", x)
}
