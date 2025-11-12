package main

import "fmt"

//Since we are using interfaces so our struck is not tightly coupled to any specific struct
// whichever struct implements the given interface can be used to create the Car struct object
type Car struct {
	EngineInterface
	TransmissionInterface
	SteeringInterface
	model string
}

func (c Car) ConvertTop() {
	fmt.Println("Converted Top...")
}