package main

import (
	"fmt"
)

func main() {
	var (
		fx float64
		x float64
	)

	fmt.Print("Masukkan nilai f(x):  ")
	fmt.Scanln(&fx)

	x = 2 / (fx + 5) + 5 

	fmt.Println("Nilai x adalah =", x)
}