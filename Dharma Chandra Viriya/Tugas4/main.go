package main

import "fmt"

func KonversiCelsiusToReamur(c float64) float64 {
	return (c * 4.0) / 5.0
}

func KonversiCelsiusToFahrenheit(c float64) float64 {
	return ((9.0 / 5.0) * c) + 32.0
}

func KonversiCelsiusToKelvin(c float64, f float64) float64 {
	return (f + 459.67) * 5 / 9
}

func main() {	
	var c float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&c)

	resultReamur := KonversiCelsiusToReamur(c)
	resultFahrenheit := KonversiCelsiusToFahrenheit(c)
	resultKelvin := KonversiCelsiusToKelvin(c, resultFahrenheit)

	fmt.Printf("Derajat Reamur: %.0f\nDerajat Fahrenheit: %.0f\nDerajat Kelvin: %.0f\n", resultReamur, resultFahrenheit, resultKelvin)
}
