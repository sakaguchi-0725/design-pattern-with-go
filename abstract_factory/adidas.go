package main

type Adidas struct{}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a *Adidas) makeShoe(size int) IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: size,
		},
	}
}

func (a *Adidas) makeShirt(size int) IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: size,
		},
	}
}
