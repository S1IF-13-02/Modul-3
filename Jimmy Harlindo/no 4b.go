package main

import "fmt"

func main() {

	var celsius float64
	
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)
	
	fahrenheit := (celsius * 9.0 / 5.0) + 32
	
	reamur := celsius * 4.0 / 5.0

	kelvin := celsius + 273.0

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}