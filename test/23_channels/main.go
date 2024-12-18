package main

import (
	"fmt"
	"time"
)

// receiving
func processNum(numChan chan int) {
	fmt.Println("processing number", <-numChan)
}

// receiving
func processNum2(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

// sending
func sum(sum chan int, num1 int, num2 int) {
	result := num1 + num2
	sum <- result
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("task processing ...")
}

// email chan only receive data
// done chan only sending data
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	// deadlock
	// messageChan := make(chan string)
	// messageChan <- "ping"

	// msg := <-messageChan

	// fmt.Println(msg)
	//deadlock end

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 5

	// numChan2 := make(chan int)

	// go processNum2(numChan2)

	// for {
	// 	numChan2 <- rand.Intn(100)
	// }

	//sending data from one goroutine to another

	// result := make(chan int)

	// go sum(result, 4, 5)

	// sum := <-result

	// fmt.Println(sum)

	// replacement of wait group
	// if there is multiple go routine then its better to use wait group, for single go routine then its better to use chan

	// done := make(chan bool)

	// go task(done)

	// <-done

	// buffer channel
	// emailChan := make(chan string, 100)

	// emailChan <- "john@example.com"
	// emailChan <- "john2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// queue channel for single sending and listening
	// emailChan := make(chan string, 100)
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("email-%d@example.com", i)
	// }
	// fmt.Println("Done sending email")

	// // its important to close the channel after sending all values
	// close(emailChan)

	// <-done

	// multiple sending and listening

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("received number:", num)
		case msg := <-chan2:
			fmt.Println("received message:", msg)
		}
	}
}
