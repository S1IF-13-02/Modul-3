package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	reamur := celsius * 4.0 / 5.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32
	kelvin := celsius + 273

	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
