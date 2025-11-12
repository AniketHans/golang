package main

type Monitor struct {
	size       int64
	resolution string
	eWasteTax  int64
	ProductDescription
}

func (m Monitor) CalculatePrice() int64 {
	return m.price + m.eWasteTax
}