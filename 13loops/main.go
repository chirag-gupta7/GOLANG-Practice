package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for d := range days {
	// 	fmt.Println(days[d])
	// }

	// for _, d := range days {
	// 	fmt.Println(d)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 3 {
			rougeValue++
			goto fahh
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println(rougeValue)
		rougeValue++
	}

fahh:
	fmt.Println("Jumped to label fahh")
}
