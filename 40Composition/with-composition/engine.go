package main

import "fmt"

// Engine
type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Engine Start...")
}

func (e Engine) Stall(){
	fmt.Println("Engine Stop...")
}