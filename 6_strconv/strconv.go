package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	b, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(b)
	}

	i, err2 := strconv.Atoi("1000")
	if err2 != nil {
		fmt.Println(err2)
	} else {

		fmt.Println(i)
	}

	fmt.Printf("strings.Split(strconv.FormatInt(999, 2), \"\"): %v\n", strings.Split(strconv.FormatInt(999, 2), ""))

	fmt.Printf("strconv.Itoa(1000): %v\n", strconv.Itoa(1000))
}
