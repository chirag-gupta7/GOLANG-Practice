package main

import "fmt"

func main() {
	defer fmt.Println("Hello, World!")
	defer fmt.Println("This will be printed first.")

	for i := 4; i > 1; i-- {
		fmt.Println("This will be printed in reverse order:", i)
	}

}
