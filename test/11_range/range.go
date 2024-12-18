package main

import "fmt"

// iterating over structures
func main() {

	// iterating arrays with range
	nums := []int{1, 2, 3, 4}

	sum := 0
	for index, num := range nums {
		fmt.Println("index", index, "num", num)
		sum += num
	}

	fmt.Println("Sum:", sum)

	// iterating maps with range

	m := map[string]string{"fName": "Jhon", "lName": "Doe"}

	// to get index and value
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	// to get only key
	for key := range m {
		fmt.Println("index", key)
	}

	// to iterate with unicode rune and starting byte of rune
	for startingByte, unicode := range "golang" {
		fmt.Println("startingByte", startingByte, "unicode", unicode, "string", string(unicode))
	}

}
