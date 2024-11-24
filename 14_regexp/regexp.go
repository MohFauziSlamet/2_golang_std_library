package main

import (
	"fmt"
	"regexp"
)

/*
Package regexp
Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
https://github.com/google/re2/wiki/Syntax
https://golang.org/pkg/regexp/
*/

func main() {
	r := regexp.MustCompile(`e([a-zA-Z])o`)

	fmt.Println(r.MatchString("edo"))
	fmt.Println(r.MatchString("ego"))
	fmt.Println(r.MatchString("eto"))
	fmt.Println(r.MatchString("eKo"))

	fmt.Printf("r.FindAllString(\"ego eto eki eko eKe eKo \", 10): %v\n", r.FindAllString("ego eto eki eko eKe eKo ", 10))
}
