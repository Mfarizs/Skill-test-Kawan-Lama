package main

import (
	"fmt"
	"math"
)

func main() {
	var hasilRuby float64
	var hasilTopaz float64
	var hasilPermata float64
	var hasil float64
	var jenis string
	var ruby = 10  // Didapatkan dari Combinasi 5
	var topaz = 6 // Didapatkan dari Combinasi 3
	var permata = 1
	var pRuby = 1000000
	var pTopaz = 1250000
	var pPermata = 3000000

	//var input string
	// var inputan float64
	// var x1, x2, x3 = 1, 1, 1
	// var kombinasi float64

	//Bila ingin Menghitung Kombinasi Menggunakan Program

	// fmt.Println("Input the jerewely Ruby/Permata/Topaz:")
	// fmt.Scan(&input)

	// if input == "Ruby" {
	// 	inputan = float64(ruby)
	// } else if input == "Topaz" {
	// 	inputan = float64(topaz)
	// } else {
	// 	inputan = float64(permata)
	// }
	// for i := 1; i <= int(inputan); i++ {
	// 	x1 = x1 * i
	// }
	// for i := 1; i <= 2; i++ {
	// 	x2 = x2 * i
	// }
	// for i := 1; i <= int(inputan-2); i++ {
	// 	x3 = x3 * i
	// }
	// kombinasi = float64(x1) / (float64(x2) * float64(x3))
	// fmt.Println(kombinasi)

	// if input == "Ruby" {
	// 	hasilRuby = float64(pRuby * int(kombinasi))
	// 	jenis = "Ruby"
	// } else if input == "Topaz" {
	// 	hasilTopaz = float64(pTopaz * int(kombinasi))
	// 	jenis = "Topaz"
	// } else {
	// 	hasilPermata = float64(pPermata * int(kombinasi))
	// 	jenis = "Permata"
	// }

	hasilRuby = float64(pRuby) * float64(ruby)
	hasilTopaz = float64(pTopaz) * float64(topaz)
	hasilPermata = float64(pPermata) * float64(permata)

	fmt.Println(hasilRuby, hasilTopaz, hasilPermata)

	//Membandingkan hasil dari harga kombinasi yang dihasilkan
	hasil = math.Max(hasilRuby, hasilTopaz)
	hasil = math.Max(hasil, hasilPermata)
	if hasil == hasilRuby {
		jenis = "Ruby"
	} else if hasil == hasilTopaz {
		jenis = "Topaz"
	} else {
		jenis = "Permata"
	}
	var kesimpulan string = fmt.Sprintf("Jumlah Keuntungan Maksimal yang didapat dengan menjual jenis %s Adalah : %v ",jenis, hasil )
	fmt.Printf(kesimpulan)
}
