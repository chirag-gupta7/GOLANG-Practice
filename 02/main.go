package main

import "fmt"

func main() {
	var username string = "Chirag"
	fmt.Println("Hello, " + username + "!")
	fmt.Printf("Type of the variable is : %T\n", username)

	var istrue bool = true
	fmt.Println(istrue)
	fmt.Printf("Type of the variable is : %T\n", istrue)

	var random uint8 = 255
	fmt.Println(random)
	fmt.Printf("Type of the variable is : %T\n", random)

	var pi float64 = 3.14
	fmt.Println(pi)
	fmt.Printf("Type of the variable is : %T\n", pi)

	var grade rune = 'A'
	fmt.Println(grade)
	fmt.Printf("Type of the variable is : %T\n", grade)

	brand := "Apple"
	fmt.Print(brand)

}
