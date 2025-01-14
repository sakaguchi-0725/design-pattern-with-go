package main

import "fmt"

type SendEmailCommand struct {
	Recipient string
	Message   string
}

func (c *SendEmailCommand) Execute() {
	fmt.Printf("Sending email to %s: %s\n", c.Recipient, c.Message)
}
