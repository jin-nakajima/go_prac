package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Whens Saturday?")
	today := time.Now().Weekday()
	fmt.Println("Today is ", today) // youbi
	fmt.Println(time.Saturday)      // Saturday
	fmt.Println(today + 3)          // Thursday
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("Too far away")
	}
}
