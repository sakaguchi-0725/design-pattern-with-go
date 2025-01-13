package main

type User struct {
	ID   uint
	Name string
}

type UserCollection struct {
	users []User
}

func (c *UserCollection) NewUserIterator() Iterator {
	return &UserIterator{
		collection: c,
		index:      0,
	}
}
