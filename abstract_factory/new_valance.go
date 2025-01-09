package main

type NewBalance struct{}

type NewBalanceShoe struct {
	Shoe
}

type NewBalanceShirt struct {
	Shirt
}

func (n *NewBalance) makeShoe(size int) IShoe {
	return &NewBalanceShoe{
		Shoe: Shoe{
			logo: "newBalance",
			size: size,
		},
	}
}

func (n *NewBalance) makeShirt(size int) IShirt {
	return &NewBalanceShirt{
		Shirt: Shirt{
			logo: "newBalance",
			size: size,
		},
	}
}
