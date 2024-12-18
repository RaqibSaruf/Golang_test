package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func getLanguages() (string, string, string, bool) {
	return "Go", "Python", "Java", true
}

func processIt(fn func(a int) int) int {
	return fn(2)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

func main() {
	resultAdd := add(1, 2)
	fmt.Println("resultAdd:", resultAdd)

	resultSub := substract(1, 2)
	fmt.Println("resultSub:", resultSub)

	lang1, lang2, lang3, err := getLanguages()
	fmt.Println("languages:", lang1, lang2, lang3, err)

	lang, _, _, err := getLanguages()
	fmt.Println("lang:", lang, err)

	fn := func(a int) int {
		return a * 2
	}

	fn2 := processIt2()
	fmt.Println(processIt(fn))
	fmt.Println(fn2(4))
}
