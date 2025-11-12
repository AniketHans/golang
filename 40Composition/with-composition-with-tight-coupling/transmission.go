package main

import "fmt"

// Transmission
type Transmission struct{}

func (t Transmission) ShiftUp(){
	fmt.Println("Gear shifted up...")
}

func (t Transmission) ShiftDown(){
	fmt.Println("Gear shifted down...")
}