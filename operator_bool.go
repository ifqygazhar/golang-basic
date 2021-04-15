package main

import "fmt"

func main() {
	nilaiLulus := 90
	absensi := 80

	nilaiAkhirLulus := nilaiLulus > 90
	nilaiAkhirAbsen := absensi > 70

	var result bool = nilaiAkhirLulus && nilaiAkhirAbsen

	fmt.Println("nilai1 = ", nilaiAkhirLulus, "nilai2 = ", nilaiAkhirAbsen, "hasil && = ", result)

}
