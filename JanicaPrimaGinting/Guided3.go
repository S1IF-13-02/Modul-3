package main 

import "fmt"

func main(){
	var rupiah, dollar int
	fmt.Scan(&rupiah)
	dollar = rupiah / 15000
	fmt.Println(dollar)
}