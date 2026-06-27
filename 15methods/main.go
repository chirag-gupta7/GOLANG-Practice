package main

import "fmt"

func main() {
	chirag := Info{
		Name:  "Chirag",
		Age:   25,
		Email: "random@go.dev",
	}
	fmt.Println(chirag)
	// fmt.Printf("Name of the user is %v and email is %v\n", chirag.Name, chirag.Email)

	chirag.getMail()
}

type Info struct {
	Name  string
	Age   int
	Email string
}

func (i Info) getMail() {
	fmt.Println("Email of the user is ", i.Email)
}
