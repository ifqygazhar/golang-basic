package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("ifqy gifha azhar", "anjay"))
	fmt.Println(strings.Trim("     ifqy gifha azhar  ", " "))
	fmt.Println(strings.Split("ifqy gifha azhar", " "))
	fmt.Println(strings.ToLower("ANJAY"))
	fmt.Println(strings.ToUpper("anjay"))
	fmt.Println(strings.ToTitle("anjay"))

}
