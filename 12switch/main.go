package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().Unix())
	randomNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is:- ", randomNumber)

	switch randomNumber {
	case 1:
		fmt.Println("Your got 1 on Dice you can move once")
	case 2:
		fmt.Println("Your got 2 on Dice you can move twice")
	case 3:
		fmt.Println("Your got 3 on Dice you can move thrice")
	case 4:
		fmt.Println("You got 4 on Dice you can move 4 times")
	case 5:
		fmt.Println("You got 5 on Dice you can move 5 times")
	case 6:
		fmt.Println("You got 6 on Dice you can move 6 times and you can roll the dice again")
		// rand.Seed(time.Now().Unix())
		randomNumber2 := rand.Intn(6) + 1
		fmt.Println("Value of second roll is:- ", randomNumber2)
	default:
		fmt.Println("What The Fuck! You got a number greater than 6 on Dice")
	}

}
