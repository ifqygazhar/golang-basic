package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
	)
	c := a + b
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)
	i++
	fmt.Println(i)
}
