package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fx := 2/(x+5) + 5
	fmt.Printf("Nilai f(x) adalah: %g\n", fx)
}
