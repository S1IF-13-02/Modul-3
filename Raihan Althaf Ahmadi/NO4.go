package main

import "fmt"
func main() {
	var celcius, fahrenheit, reamour, kelvin int
	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scan(&celcius)
	fahrenheit = (celcius + 32) * (9 / 5)
	reamour = celcius * 4 / 5
	kelvin = celcius + 273
	fmt.Println("Derajat Farenheit :" ,fahrenheit)
	fmt.Println("Derajar Reamour   :" ,reamour)
	fmt.Println("Derajat Kelvin	   :" ,kelvin)
}