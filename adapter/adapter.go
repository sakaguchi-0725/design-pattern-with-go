package main

type payPalAdapter struct {
	paypal *PayPal
}

func NewPayPalAdapter(paypal *PayPal) Payment {
	return &payPalAdapter{
		paypal: paypal,
	}
}

func (p *payPalAdapter) Pay(amount float64) {
	p.paypal.Payment(int(amount))
}
