package main

type ShippingStategy interface {
	CalculateCost(weight float64) float64
}

// 通常配送の具体的なStrategy
type RegularShipping struct{}

func (r *RegularShipping) CalculateCost(weight float64) float64 {
	return weight * 1.0 // 通常の料金計算
}

// 速達の具体的なStrategy
type ExpressShipping struct{}

func (e *ExpressShipping) CalculateCost(weight float64) float64 {
	return weight * 2.0 // 速達の料金計算
}

// 海外配送の具体的なStrategy
type InternationalShipping struct{}

func (i *InternationalShipping) CalculateCost(weight float64) float64 {
	return weight * 3.0 // 海外発送の料金計算
}
