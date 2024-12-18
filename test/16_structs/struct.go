package main

import (
	"fmt"
	"time"
)

// classic type
type Order struct {
	id        int32
	status    string
	amount    float64
	createdAt time.Time
}

// reciever type. This function is auto merged with Order type

// (*) is used to change the value of object
func (o *Order) setStatus(newStatus string) {
	o.status = newStatus
}

// no need to use (*) pointer when read value from the object
func (o Order) getAmount() float64 {
	return o.amount
}

func (o Order) print() string {
	return fmt.Sprintf("ID: %d, Status: %s, Amount: %.2f, Created At: %s", o.id, o.status, o.amount, o.createdAt.Format(time.RFC3339))
}

// constructor
func newOrder(id int32, status string, amount float64) *Order {
	order := Order{
		id:        id,
		status:    status,
		amount:    amount,
		createdAt: time.Now(),
	}

	return &order
}

func main() {
	currentOrder := Order{
		id:     1,
		status: "pending",
		amount: 100.0,
	}

	currentOrder.createdAt = time.Now()

	fmt.Println(currentOrder)
	fmt.Println(currentOrder.status)

	currentOrder2 := Order{
		id:        1,
		status:    "completed",
		amount:    100.0,
		createdAt: time.Now(),
	}

	currentOrder2.setStatus("paid")

	fmt.Println(currentOrder2)
	fmt.Println(currentOrder2.getAmount())
	fmt.Println(currentOrder2.print())

	newOrder := newOrder(3, "in progress", 500)
	fmt.Println(newOrder)

	// inline struct
	newStruct := struct {
		id   string
		name string
	}{"id", "golang"}

	fmt.Println(newStruct)
}
