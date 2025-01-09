package main

type Nike struct{}

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

func (n *Nike) makeShoe(size int) IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: size,
		},
	}
}

func (n *Nike) makeShirt(size int) IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: size,
		},
	}
}
