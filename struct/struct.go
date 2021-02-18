package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Thursday?")
	today := time.Now().Weekday()
	switch time.Thursday {
	case today + 0:
		fmt.Printf("Today")
	case today + 1:
		fmt.Printf("Tomorrow")
	case today + 2:
		fmt.Printf("In two days")
	default:
		fmt.Printf("Too far")
	}
}
