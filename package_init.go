package main

import (
	"belajar/datbase"
	"fmt"
)

func main() {
	result := datbase.GetDatabase()
	fmt.Println(result)
}
