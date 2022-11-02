package main

import "fmt"

func main() {
	const pi = 3.14

	fmt.Println("Constant PI:", pi)

	// try changin the constant, throws an error
	pi = 43
	fmt.Println("PI:", pi)
}