package main

import (
	"fmt"
)

func main() {
	var celcius float64
	fmt.Print("masukan temperatur celcius")
	fmt.Scan(&celcius)
	fahrenheit := (celcius * 9 / 5) + 32
	reamur := celcius * 4 / 5
	kelvin := celcius + 273

	fmt.Println("reamur", reamur)
	fmt.Println("fahrenheit", fahrenheit)
	fmt.Println("kelvin", kelvin)
}
