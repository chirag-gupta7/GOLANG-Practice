package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This shit is kind of difficult"

	file, err := os.Create("./logfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of file :- ", length)

	readfile("./logfile.txt")
	defer file.Close()
}

func readfile(somethingtoread string) {
	databyte, err := os.ReadFile(somethingtoread)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data in file :- ", string(databyte))

}
