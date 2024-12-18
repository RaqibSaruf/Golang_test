package main

import "fmt"

// enumurated types

type OrderStatus string

type PaymentStatus int

const (
	Pending   OrderStatus = "pending"
	Shipped               = "shipped"
	Delivered             = "delivered"
)

const (
	Unpaid PaymentStatus = iota
	Paid
	Refund
)

func printOrderStatus(status OrderStatus) {
	fmt.Println("Current order status is", status)
}

func printPaymentStatus(status PaymentStatus) {
	fmt.Println("Current order status is", status)
}

func main() {
	printOrderStatus(Shipped)
	printPaymentStatus(Paid)
}
