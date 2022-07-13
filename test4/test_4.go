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
	var ruby = 5  
	var topaz = 3 
	var permata = 1
	var pRuby = 1000000
	var pTopaz = 1250000
	var pPermata = 3000000
	var x1, x2, x3, x4, x5, x6, x7, x8, x9 = 1, 1, 1, 1, 1, 1, 1, 1, 1
	var kombinasiRuby float64
	var kombinasiTopaz float64
	var kombinasiPermata float64

	// kombinasi ruby
	for i := 1; i <= int(ruby); i++ {
		x1 = x1 * i
	}
	for i := 1; i <= 2; i++ {
		x2 = x2 * i
	}
	for i := 1; i <= int(ruby-2); i++ {
		x3 = x3 * i
	}
	kombinasiRuby = float64(x1) / (float64(x2) * float64(x3))
	fmt.Println(kombinasiRuby)

	// kombinasi Topaz
	for i := 1; i <= int(topaz); i++ {
		x4 = x4 * i
	}
	for i := 1; i <= 2; i++ {
		x5 = x5 * i
	}
	for i := 1; i <= int(topaz-2); i++ {
		x6 = x6 * i
	}
	kombinasiTopaz = float64(x4) / (float64(x5) * float64(x6))
	fmt.Println(kombinasiTopaz)

	// kombinasi Permata
	for i := 1; i <= int(permata); i++ {
		x7 = x7 * i
	}
	for i := 1; i <= 2; i++ {
		x8 = x8 * i
	}
	for i := 1; i <= int(permata-2); i++ {
		x9 = x9 * i
	}
	kombinasiPermata = float64(x7) / (float64(x8) * float64(x9))
	fmt.Println(kombinasiPermata)

	hasilRuby = float64(pRuby) * kombinasiRuby
	hasilTopaz = float64(pTopaz) * kombinasiTopaz
	hasilPermata = float64(pPermata) * kombinasiPermata

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
	var kesimpulan string = fmt.Sprintf("Jumlah Keuntungan Maksimal yang didapat dengan menjual jenis %s, dengan keuntungan : %v ", jenis, hasil)
	fmt.Printf(kesimpulan)
}
