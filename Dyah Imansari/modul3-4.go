package main

import "fmt"

func main() {
	var (
		C, F, R, K float64
	)
	fmt.Scan(&C)
	fmt.Println("Temperatur Celsius:", C)
	F = (C * 9 / 5) + 32
	fmt.Printf("Derajat Fahrenheit: %.0f\n", F)

	R = C * 4 / 5
	K = (F + 459.67) * 5 / 9
	fmt.Printf("Derajat Reamur: %.0f\n", R)
	fmt.Printf("Derajat Kelvin: %.0f\n", K)
}
