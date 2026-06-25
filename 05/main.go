package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time practice")
	currentTime := time.Now()
	fmt.Println("Current time is :-", currentTime)

	fmt.Println("Formated time is:- ", currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.April, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Created date is :-", createdDate)
	fmt.Println("Formated version of created date is :-", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
