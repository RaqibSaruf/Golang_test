package main

import (
	"fmt"
	"time"
)

type Customer struct {
	fName string
	lName string
	phone string
}

// Customer is embedded with Order. This is also called composition in go
type Order struct {
	id        int32
	status    string
	amount    float64
	createdAt time.Time
	customer  Customer
}

func main() {
	newOrder := Order{
		id:        1,
		status:    "pending",
		amount:    100.0,
		createdAt: time.Now(),
		customer: Customer{
			fName: "John",
			lName: "Doe",
			phone: "+1234567890",
		},
	}

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer.fName, newOrder.customer.lName)
}
