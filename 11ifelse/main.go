package main

import "fmt"

func main() {
	num := 7 // Change this value to test other numbers

	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}
