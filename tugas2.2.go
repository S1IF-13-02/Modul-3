package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.1415926535
	var r int
	fmt.Print("masukan jejari")
	fmt.Scan(&r)

	rFloat := float64(r)
	volumebola := (4.0 / 3.0) * pi * rFloat * rFloat * rFloat
	luasbola := 4 * pi * rFloat * rFloat
	fmt.Printf("bola dengan jejari %d memiliki volume %.4F dan luas kulit %.4F\n", r, volumebola, luasbola)
}
