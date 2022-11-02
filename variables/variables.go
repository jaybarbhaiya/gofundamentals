package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "Nigel Poulten"
	course := "Getting Started with Kubernetes"
	module := "4" // this must be an integer
	clip := 2

	fmt.Println("Name and course are set to", name, "and", course)
	fmt.Println("Module and clip are set to", module, "and", clip)

	fmt.Println("Name is TypeOf", reflect.TypeOf(name))
	fmt.Println("Module is TypeOf", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err != nil {
		fmt.Println("String to int conversion error")
	}

	total := iModule + clip
	fmt.Println("Total is", total)

	// pointers
	fmt.Println("Pointer to *course*", &course)

	var ptr *string = &course
	fmt.Println("Pointer address,", ptr, ", value of the address", *ptr)
}