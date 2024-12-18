package main

import (
	"fmt"
)

const globalValue = 100

func main() {
	const age = 35

	const (
		port = 8000
		host = "localhost"
	)

	// We can't use shorthand on constant variable declarations
	// const flag := true

	const flag = false

	fmt.Println(globalValue, age, port, host, flag)
}
