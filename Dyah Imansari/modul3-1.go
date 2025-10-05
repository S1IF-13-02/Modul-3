package main

import "fmt"

func main() {
	var fx, x float64

	fmt.Scan(&fx)
	x = (2 / (fx - 5)) - 5
	fmt.Printf("nilai x = %.0f\n", x)

}
