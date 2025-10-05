package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)
	reamur := (4.0 / 5.0) * celsius
	fahrenheit := (9.0/5.0)*celsius + 32.0
	kelvin := celsius + 273.0
	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
