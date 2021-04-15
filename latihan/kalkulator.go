package main

import "fmt"

func main() {
	fmt.Print("masukan angka pertama:")
	var a int
	fmt.Scan(&a)
	fmt.Print("masukan angka kedua :")
	var b int
	fmt.Scan(&b)

	kali := a * b
	fmt.Println("hasil a * b :", kali)
	bagi := a / b
	fmt.Println("hasil a / b :", bagi)
	sisa := a % b
	fmt.Println("hasil a % b :", sisa)
	tambah := a + b
	fmt.Println("hasil a + b :", tambah)
	kurang := a - b
	fmt.Println("hasil a - b :", kurang)

}
