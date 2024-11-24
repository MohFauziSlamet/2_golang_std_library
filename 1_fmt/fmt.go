package main

import "fmt"

func main() {
	firstName := "Muhammad Azkafitra"
	lastName := "Ramadhan"

	fmt.Println("Hello '", firstName, lastName, "'")    // -> Hello ' Muhammad Azkafitra Ramadhan ' -> dengan spasi di depan dan belakang nama
	fmt.Printf("Hello '%s %s' \n", firstName, lastName) // -> Hello 'Muhammad Azkafitra Ramadhan' -> tanpa spasi di depan dan belakang nama
}
