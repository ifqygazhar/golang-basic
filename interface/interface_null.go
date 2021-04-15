package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if 1 == 2 {
		return false
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(4)
	fmt.Println(data)
}
