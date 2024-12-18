package main

import "fmt"

type paymenter interface {
	pay(amount float64)
	refund(amount float64, account string)
}

type Payment struct {
	gateway paymenter
}

func (p Payment) makePayment(amount float64) {
	p.gateway.pay(amount)
}

func (p Payment) makeRefund(amount float64, account string) {
	p.gateway.refund(amount, account)
}

type Stripe struct{}

func (s Stripe) pay(amount float64) {
	fmt.Printf("Making a payment of %.2f using Stripe\n", amount)
}

func (s Stripe) refund(amount float64, account string) {
	fmt.Printf("Making a refund of %.2f using Stripe in %s\n", amount, account)
}

type FakePay struct{}

func (s FakePay) pay(amount float64) {
	fmt.Printf("Making a payment of %.2f using FakePay\n", amount)
}

func (s FakePay) refund(amount float64, account string) {
	fmt.Printf("Making a refund of %.2f using FakePay in %s\n", amount, account)
}

func main() {

	paymentWithStripe := Payment{
		gateway: Stripe{},
	}

	paymentWithStripe.makePayment(500)
	paymentWithStripe.makeRefund(500, "335721525")

	paymentWithFakePay := Payment{
		gateway: FakePay{},
	}

	paymentWithFakePay.makePayment(500)
	paymentWithFakePay.makeRefund(500, "335721525")
}
