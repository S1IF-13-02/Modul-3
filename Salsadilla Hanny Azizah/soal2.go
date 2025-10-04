package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		jari float64
		volume float64
		luasbola float64
	)

	fmt.Print("Jejari =  ")
	fmt.Scanln(&jari)

	volume = (4.0 / 3.0) * math.Pi * jari * jari * jari
	luasbola = 4.0 * math.Pi * jari * jari
	
	fmt.Println("Bola dengan jejari", jari, "memiliki volume", volume, "dan luas kulit", luasbola)
}