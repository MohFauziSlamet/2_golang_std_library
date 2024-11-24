package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CreateNewFile(name string , massage string) error  {
//* open file 
os.OpenFile(name, os.o)
}

func main() {
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()

	

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()

	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("Hello world \n")
	writer.WriteString("Selamat belajar")

	writer.Flush()

	fmt.Println()
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()
}
