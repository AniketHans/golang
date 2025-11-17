package main

import "fmt"

type BillableProduct interface {
	// Whichever type/struct has functions with the below signatures will implement the BillableProduct interface
	CalculatePrice() int64
}

func CartTotal(products ...BillableProduct) int64 {
	// The products can have different logic in CalculatePrice() but as soon as they implements the BillableProduct interface,
	// we can loop over them and call the CalculatePrice(). 
	totPrice := int64(0)
	for _, product := range products {
		//This is polymorphism as the implementation of the CalculatePrice() by each product can be different but since they have
		// a common function signature, they implements the 
		totPrice += product.CalculatePrice()
	}
	return totPrice
}

func main() {
	s := Shirt{42, "Blue", 100, ProductDescription{400, "PeterEngland"}}
	m := Monitor{size: 52, resolution: "4K", eWasteTax: 1000, ProductDescription: ProductDescription{price: 10000, brand: "Dell"}}
	w := Whine{year: 2000, qty: "145 ml", alcoholTax: 900, ProductDescription: ProductDescription{2000, "Desi"}}

	totalCartValue := CartTotal(s, m, w)
	fmt.Println("Cart total is",totalCartValue)
}