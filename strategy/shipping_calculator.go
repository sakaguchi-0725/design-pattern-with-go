package main

type ShippingCostCalculator struct {
	strategy ShippingStategy
}

func (s *ShippingCostCalculator) SetStrategy(strategy ShippingStategy) {
	s.strategy = strategy
}

func (s *ShippingCostCalculator) Calculate(weight float64) float64 {
	if s.strategy == nil {
		panic("Shipping strategy is not set")
	}
	return s.strategy.CalculateCost(weight)
}
