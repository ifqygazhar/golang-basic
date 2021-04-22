// S = jarak
// t = waktu tempuh (jam)
// v = kecepatan jam (km/jam)

package main

import (
	"fmt"
)

func main() {
	var v float64

	fmt.Print("Masukan jarak = ")
	var s int
	fmt.Scan(&s)

	fmt.Print("Masukan waktu tempuh = ")
	var t int
	fmt.Scan(&t)
	v = float64(s) / float64(t)
	// f := float64(v)

	fmt.Println("Diketahui S = ", s, " dan T = ", t,"jam",",Ditanya kecepatan rata - rata (v)")
	fmt.Println("V = S/t = ", s, "/", t)
	fmt.Println("V = ", v)
	fmt.Println("Jadi kecepatan rata-rata adalah = ", v,"km/jam")
}
