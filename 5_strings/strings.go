package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Azkafitra Ramadhan", "Ramadhan"))
	fmt.Println(strings.Split("Azkafitra Ramadhan", " "))
	fmt.Println(strings.ToLower("Azkafitra Ramadhan"))
	fmt.Println(strings.ToUpper("Azkafitra Ramadhan"))
	fmt.Println(strings.Trim("				Azkafitra Ramadhan			", "	"))
	fmt.Println(strings.ReplaceAll("Azkafitra Ramadhan", "Ramadhan", "Avril"))
}
