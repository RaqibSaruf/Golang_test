package main

import (
	"fmt"
)

func main() {
	var text1 string = "Hello world"
	text2 := "Hello world"
	fmt.Println(text1, text2)
	var number1 int16 = 123
	number2 := 123
	var number3 int64 = 123
	fmt.Println(number1, number2, number3)
	var isNotValid bool = false
	isValid := true
	fmt.Println(isValid, isNotValid)
	var floatNumber1 float32 = 50.5
	floatNumber2 := 50.5
	fmt.Println(floatNumber1, floatNumber2)

}
