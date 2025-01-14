package main

import "fmt"

type SendLineCommand struct {
	LineID  string
	Message string
}

func (c *SendLineCommand) Execute() {
	fmt.Printf("Sending LINE to %s: %s\n", c.LineID, c.Message)
}
