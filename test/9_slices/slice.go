package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slice is nil
	var numbers []int

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(numbers == nil)

	// type -> []int
	// size -> 2 -> optional
	// capacity -> 5 -> optional
	var numbers2 = make([]int, 2, 5)
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	numbers2 = append(numbers2, 3)
	numbers2 = append(numbers2, 4)
	numbers2 = append(numbers2, 5)

	fmt.Println(numbers2)
	fmt.Println(cap(numbers2))

	var numbers3 = make([]int, 0, 5)
	numbers3 = append(numbers3, 1)
	numbers3 = append(numbers3, 2)
	numbers3 = append(numbers3, 3)

	fmt.Println(numbers3)
	fmt.Println(cap(numbers3))

	var numbers4 = make([]int, 2)
	numbers4 = append(numbers4, 23)
	numbers4 = append(numbers4, 25)

	var numbers5 = make([]int, len(numbers4))

	copy(numbers5, numbers4)

	fmt.Println(numbers4, numbers5)

	//slice operators
	var numbers6 = []int{1, 2, 3, 4}

	// start index to end index
	fmt.Println("the value of start index to end index", numbers6[1:2])
	fmt.Println("if start index is not available then set zero'th value in the index", numbers6[:2])

	//slice built in method
	numbers7 := []int{1, 2, 3, 4}
	fmt.Println(numbers7)

	numbers8 := []int{1, 2, 3, 4}
	fmt.Println(numbers8)

	fmt.Println(slices.Equal(numbers7, numbers8))

}
