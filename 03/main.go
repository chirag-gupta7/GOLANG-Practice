package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	fmt.Printf("Your name is %v and your age is %v", name, age)

}
