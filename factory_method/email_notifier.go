package main

import "fmt"

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(msg string) {
	fmt.Println("メール送信: ", msg)
}

type EmailNotifierFactory struct{}

func (e *EmailNotifierFactory) NewNotifier() Notifier {
	return &EmailNotifier{}
}
