package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}

	valueInt, _ := strconv.Atoi("10000")
	fmt.Println(valueInt)
}
