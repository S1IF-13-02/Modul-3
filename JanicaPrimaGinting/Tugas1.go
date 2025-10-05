package main

import "fmt"

func main() {
	var fx float64
	fmt.Scan(&fx)
	x := 2/(fx-5) - 5
	fmt.Printf("%.1g\n", x)
}
