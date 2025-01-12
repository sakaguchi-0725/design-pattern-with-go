package main

import "fmt"

// Adaptee
type PayPal struct{}

func (p *PayPal) Payment(amount int) {
	fmt.Println("Payment amount for PayPal: ", amount)
}
