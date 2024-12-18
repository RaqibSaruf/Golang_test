package main

import "fmt"

func main() {
	// if else
	age := 25
	if age > 18 {
		fmt.Println(`Adult`)
	} else {
		fmt.Println(`Child`)
	}

	// else if
	age = 15
	if age > 18 {
		fmt.Println(`Adult`)
	} else if age > 12 {
		fmt.Println(`Teenager`)
	} else {
		fmt.Println(`Child`)
	}

	// shorthand if else
	if age := 10; age > 18 {
		fmt.Println(`Adult`)
	} else if age > 12 {
		fmt.Println(`Teenager`)
	} else {
		fmt.Println(`Child`)
	}

	// go does not have any ternary operator, we have to use always normal if else

}
