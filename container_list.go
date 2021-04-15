package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("ifqy")
	data.PushBack("gifha")
	data.PushBack("azhar")

	// belakang ke depan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("================")
	//belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
