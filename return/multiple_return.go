package main

import "fmt"

func getName() (string, int) {
	return "ifqy", 16
}

func main() {
	firstname, _ := getName()
	fmt.Println(firstname)
}
