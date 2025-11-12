package main

type Whine struct {
	year       int64
	qty        string
	alcoholTax int64
	ProductDescription
}

func (w Whine) CalculatePrice() int64 {
	return w.price + w.alcoholTax
}