package main

import "fmt"

func main() {
	var array = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(array)

	var array2 [3]string

	array2[0] = "ifqy"
	array2[1] = "gifha"
	array2[2] = "azhar"

	fmt.Println(array2)
	fmt.Println(len(array2))
	array[0] = 4
	fmt.Print(array[0])
}
