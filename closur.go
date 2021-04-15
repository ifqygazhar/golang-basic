package main

import "fmt"

func main() {
	name := "ifqy"
	counter := 0

	// increment := func() {
	// 	name = "anjay"
	// 	fmt.Println("increment") //salah
	// 	counter++
	// }

	benar := func() {
		var name string
		fmt.Println("masukan nama = ")
		_, cek := fmt.Scanf("%d", &name)
		fmt.Println("benar") //benar
		fmt.Println(cek)

	}

	benar()
	fmt.Println(name)
	fmt.Println(counter)
}
