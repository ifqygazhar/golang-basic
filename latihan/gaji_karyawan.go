package main

import "fmt"

func main() {
	// Menghitung gaji karyawan plus bonus lembur
	var upah int
	fmt.Print("masukan nama anda = ")
	var nama string
	fmt.Scan(&nama)

	fmt.Print("masukan golongan gaji anda = ")
	var golongan string
	fmt.Scan(&golongan)

	fmt.Print("masukan jam kerja gaji anda = ")
	var jam int
	fmt.Scan(&jam)

	switch golongan {
	case "A":
		upah = 5000
	case "B":
		upah = 7000
	case "C":
		upah = 8000
	case "D":
		upah = 10000

	default:
		fmt.Println("tidak ada dalam kategori")
	}

	total_upah := jam * upah

	if jam-48 > 0 {
		total_upah = total_upah + ((jam - 48) * 4000)
	}

	fmt.Println(nama, "menerima upah Rp.", total_upah, "per minggu nya")
}
