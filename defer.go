package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil")
}

func runApp(value int) {
	defer logging()
	fmt.Println("run aplikasi")
	result := 10 / value
	fmt.Println(result)

}

func main() {
	runApp(10)
}
