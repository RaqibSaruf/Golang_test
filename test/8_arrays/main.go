package main

import "fmt"

// numbered sequence of specific length like C/C++
func main() {
	// add default zeros value of the type assigned when array is declared and not added any value onto it
	var numbers [3]int

	// print length
	fmt.Println(len(numbers))

	numbers[0] = 56
	fmt.Println(numbers[0])

	numbers = [3]int{1, 2, 3}

	var animals = [3]string{"dog", "cat", "cow"}
	fmt.Println(animals)

	// 2d array

	twoDimensionNumbers := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDimensionNumbers)
}
