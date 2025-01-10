package main

type Notifier interface {
	Notify(msg string)
}

type NotifierFactory interface {
	NewNotifier() Notifier
}
