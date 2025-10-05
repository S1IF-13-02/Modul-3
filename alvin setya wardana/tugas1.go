package main

import "fmt"

func main() {

	var x float64
	var Fx float64

	fmt.Print("Masukan nilai F(x): ")
	fmt.Scanln(&x)
	Fx = (2 / (x + 5)) + 5
	fmt.Printf("nilai x adalah %.2f/", Fx)

	konversi1 := int(Fx)
	fmt.Println(" konversi nilai float ke integer =", konversi1)

}
