package main

import "fmt"

type Car struct{}

type Truck struct{}

// Car functionality

func (c Car) ShiftUp() {
	fmt.Println("Gear shifted up...")
}

func (c Car) Start() {
	fmt.Println("Engine started...")
}

func (c Car) TurnLeft() {
	fmt.Println("Steared Left...")
}

func (c Car) ConvertTop() {
	fmt.Println("Coverted Top...")
}

// Truck Functionality
func (t Truck) ShiftUp() {
	fmt.Println("Gear shifted up...")
}

func (t Truck) Start() {
	fmt.Println("Engine started...")
}

func (t Truck) TurnLeft() {
	fmt.Println("Steared Left...")
}

func (t Truck) FourWheelDrive() {
	fmt.Println("4WD engaged...")
}

func main(){
	c := Car{}
	t := Truck{}

	c.Start()
	c.TurnLeft()
	c.ConvertTop()

	t.Start()
	t.ShiftUp()
	t.FourWheelDrive()
}