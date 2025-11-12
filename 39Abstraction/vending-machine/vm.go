package vm

import "fmt"

type VendinMachineSupertech struct{}

func New() *VendinMachineSupertech {
	return &VendinMachineSupertech{}
}

func (vm *VendinMachineSupertech) GetDrink(money int64, drink string) string {
	return fmt.Sprintf("Ice cold %v",drink)
}