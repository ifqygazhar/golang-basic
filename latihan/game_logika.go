package main

import "fmt"

func main() {
	fmt.Println("pilih 1 = math, pilih 2 = logika")
	fmt.Print("masukan pilihan 1 atau 2 = ")
	var pilih1 string
	fmt.Scan(&pilih1)
	if pilih1 == "1" {
		fmt.Println("anda berada di zona matematika jawab soal ini")
		fmt.Print("5 * 5 = ")
		var kali int
		fmt.Scan(&kali)

		if kali != 25 {
			fmt.Print("salah")
		}

	} else if pilih1 == "2" {
		fmt.Println("anda berada di zona teka teki jawab soal ini")
		fmt.Print("jika hari ini rabu maka 2 hari kebelakang nya adalah = ")
		var hasil string
		fmt.Scan(&hasil)

		if hasil != "senin" {
			fmt.Print("salah")
		}
	} else {
		fmt.Println(pilih1, "tidak dalam kategori")
	}
}
