package main

import "flag"

/*
	Package flag
	Package flag berisikan fungsionalitas untuk memparsing command line argument
	jadi mengambil argument dengan pacakge OS , lalu memparsing nya dengan package FLAG
	https://golang.org/pkg/flag/
*/

func main() {
	/*
		flag.String() :
		String defines a string flag with specified name, default value, and usage string.
		The return value is the address of a string variable that stores the value of the flag.

		flag.Int() :
		Int defines a Int flag with specified name, default value, and usage Int.
		The return value is the address of a Int variable that stores the value of the flag.
	*/

	var username *string = flag.String("username", "root", "username database")
	var password *string = flag.String("password", "root", "password database")
	var host *string = flag.String("host", "root", "host database")
	var port *int = flag.Int("port", 0, "port database")

	//* jangan lupa memanggil method flag.Parse()
	flag.Parse()

	println("Username", *username)
	println("Password", *password)
	println("Host", *host)
	println("Port", *port)
}
