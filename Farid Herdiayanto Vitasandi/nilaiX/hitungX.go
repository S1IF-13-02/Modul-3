package main

import (
	"fmt"
	"math"
)

func main() {
	
	var fx float64

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&fx)

	x := (2 / (fx - 5)) - 5

	fmt.Println("Nilai x adalah: ", int(math.Round(x)))
}
