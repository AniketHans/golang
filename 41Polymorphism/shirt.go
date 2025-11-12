package main

type Shirt struct {
	size     int64
	color    string
	discount int64
	ProductDescription
}

func (s Shirt) CalculatePrice() int64 {
	return s.price - s.discount
}