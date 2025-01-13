package main

type Iterator interface {
	HasNext() bool
	Next() User
}

// UserIterator UserCollection専用のIterator
type UserIterator struct {
	collection *UserCollection
	index      int
}

func (i *UserIterator) HasNext() bool {
	return i.index < len(i.collection.users)
}

func (i *UserIterator) Next() User {
	if !i.HasNext() {
		panic("No more users")
	}
	user := i.collection.users[i.index]
	i.index++
	return user
}
