package main

import "fmt"

func main() {
	var r , v , l float64
	const pi = 3.1415926535

	fmt.Print("Masukan jejari :")
	fmt.Scan(&r)

	v = 4.0/3.0*pi*r*r*r
	l = 4*pi*r*r

	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f",r,v,l)
}
