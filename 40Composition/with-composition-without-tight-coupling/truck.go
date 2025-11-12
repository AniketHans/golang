package main

import "fmt"

//Since we are using interfaces so our struck is not tightly coupled to any specific struct
// whichever struct implements the given interface can be used to create the Truck struct object
type Truck struct {
	EngineInterface
	TransmissionInterface
	SteeringInterface
	model string
}

func (t Truck) FourWheelDrive() {
	fmt.Println("4WD Engaged")
}