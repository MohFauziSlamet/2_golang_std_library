package main

import (
	"container/list"
	"fmt"
)

/*
Package container/list :
Package container/list adalah implementasi struktur data double linked list di Go-Lang.

double linked list biasanya di gunakan untuk mengolah data antrian.
https://golang.org/pkg/container/list/
*/
func main() {
	e := list.New()

	e.PushBack("Muhammad")
	e.PushBack("Azkafitra")
	e.PushBack("Ramadhan")

	head := e.Front()

	fmt.Println(head.Value) // Muhammad

	next := head.Next()
	fmt.Println(next.Value) // Azkafitra

	next = next.Next()
	fmt.Println(next.Value) // Ramadhan

	fmt.Println("========================")

	//* looping data linked list
	for data := e.Front(); data != nil; data = data.Next() {
		fmt.Println(data.Value)
	}
}
