package main

import "fmt"

func faktorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result = result * i
	}
	return result
}

func faktorialRekursip(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorialRekursip(value-1)
	}
}

func main() {
	loop := faktorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	rekursip := faktorialRekursip(5)
	fmt.Println(rekursip)
}
