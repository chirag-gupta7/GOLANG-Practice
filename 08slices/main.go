package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3, 4, 5)

	fmt.Println("Slice elements:", slice)

	new_slice := make([]string, 2)
	new_slice[0] = "Hello"
	new_slice[1] = "World"

	new_slice = append(new_slice, "from", "Go")
	fmt.Println("Slice using make() :- ", new_slice)

	cources := []string{"C", "C++", "Java", "Python"}
	var index int = 1
	fmt.Println(cources)
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)
}
