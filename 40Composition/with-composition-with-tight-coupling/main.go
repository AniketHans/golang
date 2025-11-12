package main

import "fmt"

type Cars struct {
	// Using composition below, Cars struct gets the functionalities of Engine, Transmission and Steering
	Engine
	Transmission
	Steering
	model string
}

func (c Cars) ConvertTop() {
	fmt.Println("Top Converted...")
}

type Trucks struct{
	// Using composition below, Trucks struct gets the functionalities of Engine, Transmission and Steering
	Engine
	Transmission
	Steering
	model string
}

func (t Trucks) FourWheelDrive() {
	fmt.Println("4WD engaged...")
}

func main(){
	c := Cars{Engine{},Transmission{},Steering{},"Ferrari"}
	t := Trucks{Engine{}, Transmission{}, Steering{}, "Tata"}

	c.Start()
	c.Stall()
	c.ConvertTop()

	t.TurnLeft()
	t.Start()
	t.FourWheelDrive()

	StartEngines(c,t)
}

type Startables interface{
	// All the types/structs having the start functionality implements Startables interface
	Start()
}

func StartEngines(vehicles ...Startables){
	// Function to start engines of all the vehicles which implements the Startable Interafce
	for vehicleNo := range vehicles{
		vehicles[vehicleNo].Start()
	} 
}

// Since the package main is distributed amoung multiple files use `go run .` to run the code

// In this code, Cars and Trucks structs are tightly coupled to specific Engine, Transmission and Steering structs
// If there is say another Transmission say AdvancedTransmission having the same functions as the Transmission, we cannot use it without
// making change to the Cars and Trucks structs