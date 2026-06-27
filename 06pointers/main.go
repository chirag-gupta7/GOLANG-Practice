package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's practice Pointers")
	var a int = 42
	var ptr *int = &a
	fmt.Println("Values of a :", *ptr)
	// fmt.Println(*a)

}
