package main

import "fmt"

type TransmissionInterface interface{
	ShiftUp()
	ShiftDown()
}

// Transmission
type Transmission struct{}

func (t Transmission) ShiftUp(){
	fmt.Println("Gear shifted up...")
}

func (t Transmission) ShiftDown(){
	fmt.Println("Gear shifted down...")
}


// AdvancedTransmission

type AdvancedTransmission struct{}

func (at AdvancedTransmission) ShiftUp(){
	fmt.Println("Quick Gear shifter up...")
}

func (at AdvancedTransmission) ShiftDown(){
	fmt.Println("Quick Gear shifter down...")
}