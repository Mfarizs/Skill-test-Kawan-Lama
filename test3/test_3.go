package main

import (
	"fmt"
)

func main() {
	var hasil int
	var j int
	var k int
	fmt.Println("Input the number of J :")
	fmt.Scan(&j)
	fmt.Println("Input the number of K :")
	fmt.Scan(&k)
	for j > 0 {
		hasil += k
		j--
	}
	fmt.Println("Hasil Perkalian Adalah :", hasil)

}