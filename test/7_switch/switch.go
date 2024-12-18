package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch

	i := 1

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Unknown")
	}

	// multiple conditions

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	// type switch

	getType := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("%v is an %T\n", i, t)
		case string:
			fmt.Printf("%v is a %T\n", i, t)
		default:
			fmt.Printf("%v is an unknown %T\n", i, t)
		}
	}

	getType("golang")
	getType(255)
	getType(29.56)
}
