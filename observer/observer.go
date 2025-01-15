package main

import "fmt"

type Observer interface {
	Update(news string)
}

type User struct {
	Name string
}

func (u *User) Update(news string) {
	fmt.Printf("User %s received news: %s\n", u.Name, news)
}
