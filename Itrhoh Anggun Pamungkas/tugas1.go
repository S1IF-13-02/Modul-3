package main

import (
	"fmt"
)

func main() {
	var fx float64

	fmt.Print("nilai f(x): ")
	fmt.Scan(&fx)

	x := (2 / (fx - 5)) - 5
	fmt.Printf("Nilai x = %.0f\n", x)
}

// f(x) = 2/(x+5) + 5
// -> x = (2 / (f(x) - 5)) - 5
