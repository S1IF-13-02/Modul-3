package main

import "fmt"

func main() {
	var celcius float64
	fmt.Print("Temperature celcius: ")
	fmt.Scan(&celcius)

	fahrenheit := (celcius * 9.0 / 5.0) + 32
	reamur:= celcius * 4.0 / 5.0
	kelvin := celcius + 273.15

	fmt.Printf("Derajat reamur: %.0f\n", reamur)
	fmt.Printf("Derajat fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat kelvin: %.0f \n", kelvin)
}
