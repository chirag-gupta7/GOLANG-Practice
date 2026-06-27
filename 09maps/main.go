package main

import "fmt"

func main() {
	newmap := make(map[int]string)

	newmap[1] = "I"
	newmap[2] = "II"
	newmap[3] = "III"
	newmap[4] = "IV"
	newmap[5] = "V"

	for key, value := range newmap {
		fmt.Printf("For the numbre %v roman numeral is %v\n", key, value)
	}

}
