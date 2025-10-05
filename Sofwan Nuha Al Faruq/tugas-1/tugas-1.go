package main

import "fmt"

func main() {
	var fx float64
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	nilaiX :=(2/(fx-5)) - 5
	fmt.Printf("Nilai X nya adalah %.0f", nilaiX)
}