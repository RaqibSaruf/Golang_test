package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
		go printSomething("jhon" + string(i))
	}

	time.Sleep(time.Second * 2)
}
