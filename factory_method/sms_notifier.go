package main

import "fmt"

type SMSNotifier struct{}

func (s *SMSNotifier) Notify(msg string) {
	fmt.Println("SMS送信: ", msg)
}

type SMSNotifierFactory struct{}

func (s *SMSNotifierFactory) NewNotifier() Notifier {
	return &SMSNotifier{}
}
