package main

import "fmt"

// For is the only construct for looping in golang
func main() {

	// while loop
	var i = 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		} // break the loop if condition is met
	}

	// classic for loop

	for i = 0; i <= 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// range loop

	i = 0.
	for i := range 5 {
		fmt.Println(i)
		i++
	}

}
