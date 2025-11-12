package main

import "fmt"

type SteeringInterface interface{
	TurnLeft()
	TurnRight()
}

// Steering
type Steering struct{}

func (s Steering) TurnLeft(){
	fmt.Println("Turned Left...")
}

func (s Steering) TurnRight(){
	fmt.Println("Turned Right...")
}
