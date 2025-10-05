package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	fx := (2 / (x + 5)) + 5
	fx = float64(int(fx + 0.5))
	fmt.Printf("%.0f\n", fx)

}
