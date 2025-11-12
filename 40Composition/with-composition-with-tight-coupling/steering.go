package main

import "fmt"

// Steering
type Steering struct{}

func (s Steering) TurnLeft(){
	fmt.Println("Turned Left...")
}

func (s Steering) TurnRight(){
	fmt.Println("Turned Right...")
}
