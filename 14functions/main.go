package main

import "fmt"

func main() {
	greet()

	grander := adder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Sum of first 10 positive integer : ", grander)
}

func greet() {
	fmt.Println("Hello Nigga!")
}

func adder(x ...int) int {
	total := 0

	for _, i := range x {
		total += i

	}
	return total
}
