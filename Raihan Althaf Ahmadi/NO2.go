package main

import (
	"fmt"
)

func main() {
	var r float32
	var luas float32
	var volume float32
	var pi float32

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	pi = 3.1415926535
	volume = (4.0 / 3.0) * pi * r * r * r
	luas = 4 * pi * r * r

	fmt.Println("Bola dengan jejari", r ," memiliki volume", volume, " dan luas kulit ", luas)
}
