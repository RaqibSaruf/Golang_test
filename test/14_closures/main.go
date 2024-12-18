package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}
func main() {
	increament := counter()
	fmt.Println(increament()) // 1
	fmt.Println(increament()) // 1
}
