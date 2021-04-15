package main

import "fmt"

type Man struct {
	name, name2, name3 string
}

func (man *Man) Married() {
	man.name = "Mr." + man.name
	man.name2 = "Mrs." + man.name2
	man.name3 = "Drs." + man.name3
}

func main() {
	az := Man{"azhar", "mariam", "ifqy"}
	az.Married()

	fmt.Println(az)
}
