package main

import (
	vm "abstraction/vending-machine"
	"fmt"
)

type VendingMachine interface {
	GetDrink(money int64, drink string) string
}

type DrinksApplication struct {
	vm VendingMachine
}

func NewApplication(vm VendingMachine) *DrinksApplication {
	return &DrinksApplication{vm: vm}
}

func main() {
	machine := vm.New()
	app := NewApplication(machine)
	fmt.Println("Drink received :-",app.vm.GetDrink(100, "Cola"))
}