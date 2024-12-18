package main

import "fmt"

func changeByValue(num int) {
	num = 5
	fmt.Println("In change by value", num)
}

func changeByReferenc(num *int) {
	*num = 5
	fmt.Println("In change by reference", *num)
}
func main() {
	num := 2
	changeByValue(num)
	fmt.Println("Number after change value", num)

	changeByReferenc(&num)
	fmt.Println("Number after change reference", num)
}
