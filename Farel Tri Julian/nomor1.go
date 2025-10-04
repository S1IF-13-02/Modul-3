package main

import "fmt"

func main() {

	var fx float64

	fmt.Scanln(&fx)

	x := (2 / (fx + 5)) + 5

	fmt.Printf("%.0f\n", x)
}
