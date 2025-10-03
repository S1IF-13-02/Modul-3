package main 

import "fmt"

func main() {
	var celcius, fahrenheit, reamur, kelvin float64

	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)

	fahrenheit = (9.0 / 5.0) * celcius + 32
	reamur = (4.0 / 5.0) * celcius
	kelvin = celcius + 273.15

	fmt.Print("\nTemperatur dalam Celcius: ", celcius, "°C\n")
	fmt.Printf("Temperatur dalam Fahrenheit: %.2f °F\n", fahrenheit)
	fmt.Printf("Temperatur dalam Reamur: %.2f °R\n", reamur)
	fmt.Print("Temperatur dalam Kelvin: ", kelvin, "K\n")
}
