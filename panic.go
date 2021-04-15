package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("apk eror dengan kode :", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("apk error")
	}
	fmt.Println("apk jalan")
}

func main() {
	runApp(false)
}
