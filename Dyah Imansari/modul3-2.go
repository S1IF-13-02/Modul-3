package main

import "fmt"

func main() {
	var (
		r, V, L float64
		π       = 3.1415926535
	)
	fmt.Scan(&r)
	V = (4.0 / 3.0) * π * r * r * r
	L = 4 * π * r * r
	fmt.Println("Jejari =", r)
	fmt.Println("Bola dengan jejari", r)
	fmt.Printf("memiliki volume %.4f\n", V)
	fmt.Printf("dan luas kulit %.4f\n", L)
}
