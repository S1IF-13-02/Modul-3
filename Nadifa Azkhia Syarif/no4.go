package main

import (
	"fmt"
)

func main() {
	// langsung pake nilai default
	celsius := 50.0

	// konversi
	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273

	// output
	fmt.Println("Temperatur Celsius:", celsius)
	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
