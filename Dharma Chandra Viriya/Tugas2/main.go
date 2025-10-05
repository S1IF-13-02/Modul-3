package main

import (
	"fmt"
	"math"
)

func HitungVolumeBola(r float64, pi float64) float64 {
	return (4 * (pi * (r * r * r))) / 3
}

func HitungLuasBola(r float64, pi float64) float64 {
	return 4 * pi * (r * r)
}

func main() {
	const pi float64 = math.Pi
	var r int

	fmt.Print("Masukkan r: ")
	fmt.Scan(&r)

	resultVolumeBola := HitungVolumeBola(float64(r), pi)
	resultLuasBola := HitungLuasBola(float64(r), pi)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, resultVolumeBola, resultLuasBola)
}
