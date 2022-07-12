package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var hasil int
	var random string
	var s = 0
	var is = 0
	var a = 0
	var p = 0

	fmt.Println("Input the Random String  :")
	fmt.Scan(&random)
	var dr = strings.Split(random, "")
	log.Println(dr)
	var as = strings.Split("SIAPA", "")
	fmt.Println(as)

	for i := 0; i < len(dr); i++ {
		for j := 0; j < len(as); j++ {
			if strings.Compare(dr[i], as[j]) == 0 {
				if strings.Compare(as[j], "S") == 0 {
					s++
				} else if strings.Compare(as[j], "I") == 0 {
					is++
				} else if strings.Compare(as[j], "A") == 0 {
					a++
				} else if strings.Compare(as[j], "P") == 0 {
					p++
				}
			}
		}
		for s != 0 && is != 0 && a != 0 && p != 0 {
			if s >= 1 && is >= 1 && a >= 2 && p >= 1 {
				s--
				is--
				a--
				p--
				hasil++
			}
		}

	}

	fmt.Println("Jumlah Kata yang dibuat adalah :", hasil)

}