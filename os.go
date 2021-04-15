package main

import (
	"fmt"
	"os"
)

func main() {
	result := os.Args
	fmt.Println(result)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println("error", err.Error())
	}
}
