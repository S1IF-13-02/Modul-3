package main

import (
	"fmt"
)

func main() {
	var (
		fahrenheit float64
		celcius float64
		reamur float64
		kelvin float64
	)

	fmt.Print("Temperatur Celcius =  ")
	fmt.Scanln(&celcius)

	fahrenheit = (celcius * 9.0 / 5.0) + 32
	reamur =  celcius * 4.0 / 5.0 
	kelvin = (fahrenheit + 459.67) * 5.0 / 9.0

	
	fmt.Println("Derajat Reamur = ", reamur)
	fmt.Println("Derajat Fahrenheit = ", fahrenheit)
	fmt.Println("Derajat Kelvin =", kelvin)
}