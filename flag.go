package main

import (
	"flag"
	"fmt"
)

func main() {
	var hostname *string = flag.String("host", "localhost", "put your datbase host")
	var user *string = flag.String("user", "root", "put your database user")
	var password *string = flag.String("pasword", "root", "put your database password")

	flag.Parse()

	fmt.Println("host:", *hostname)
	fmt.Println("user:", *user)
	fmt.Println("pass:", *password)
}
