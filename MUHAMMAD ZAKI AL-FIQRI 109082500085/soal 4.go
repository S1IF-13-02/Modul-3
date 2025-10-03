package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celcius)

	reamur := (celcius * (4.0 / 5.0))
	fahrenheit := (celcius*(9.0/5.0) + 32)
	kelvin := celcius + 273

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
