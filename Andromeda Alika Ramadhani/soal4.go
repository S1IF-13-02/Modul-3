package main

import "fmt"

func main() {
	var celcius int
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&celcius)

	Derajat_Fahreinheit := (celcius - 32) * 5 / 9
	Derajat_Reamur := celcius * 4 / 5
	Derajat_Kelvin := celcius + 273

	fmt.Printf("Derajat Fahreinheit: %d\n", Derajat_Fahreinheit)
	fmt.Printf("Derajat Reamur: %d\n", Derajat_Reamur)
	fmt.Printf("Derajat Kelvin: %d\n", Derajat_Kelvin)
}
