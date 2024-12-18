package main

import "fmt"

// accept any variable , we can also use interface{} instead of any
func printSomething[T any](printable []T) {
	fmt.Println(printable)
}

// accept only int or string
func printSomeIntOrString[T int | string](printable []T) {
	fmt.Println(printable)
}

type Stack[T any] struct {
	items []T
}

// we can use comparable instead of interface{} or any
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strings := []string{"golang", "js"}
	booleans := []bool{true, false, false}
	printSomething(nums)
	printSomething(strings)
	printSomething(booleans)
	printSomeIntOrString(nums)
	printSomeIntOrString(strings)

	mystack := Stack[string]{
		items: []string{"apple", "banana", "cherry"},
	}

	fmt.Println(mystack.items)

	vals := []bool{true, false, true}

	printSlice(vals, "Jhon")
}
