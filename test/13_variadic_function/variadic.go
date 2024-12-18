package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func processIt(a ...interface{}) (int, []interface{}) {
	total := 0
	arrays := make([]interface{}, 0, 5)
	for _, v := range a {
		switch v := v.(type) {
		case int:
			fmt.Println("integer", v)
			total = total + v
		default:
			fmt.Println("interface", v)
			arrays = append(arrays, v)
		}
	}

	return total, arrays
}

func main() {
	total := sum(1, 5, 8, 74, 3)

	fmt.Println(total)

	nums := []int{6, 4, 8, 6}

	total = sum(nums...)

	fmt.Println(total)

	total, arrays := processIt(1, 2, 3, 4, "hello", 5.5, []int{6, 7})
	fmt.Println(total)
	fmt.Println(arrays)
}
