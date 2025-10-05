package main

import (
	"fmt"
)

func main() {
	var fx, x float64
	fmt.Print("masukan nilai fx")
	fmt.Scan(&fx)

	x = (2 / (fx - 5)) - 5

	fmt.Printf("hasil nilai x: %.0F\n", x)

}
