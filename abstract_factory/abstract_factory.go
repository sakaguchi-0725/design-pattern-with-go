package main

import "fmt"

type Brand string

const (
	adidas     Brand = "adidas"
	nike       Brand = "nike"
	newValance Brand = "newBalance"
)

type ISportsFactory interface {
	makeShoe(size int) IShoe
	makeShirt(size int) IShirt
}

func GetSportsFactory(brand Brand) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	case "newBalance":
		return &NewBalance{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
