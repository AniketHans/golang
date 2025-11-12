package main

import "fmt"


type EngineInterface interface{
	Start()
	Stall()
}

// Engine
type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Engine Start...")
}

func (e Engine) Stall(){
	fmt.Println("Engine Stop...")
}

//Turbo Engine

type TurboEngine struct{}

func (te TurboEngine) Start(){
	fmt.Println("Turbo Engine Start...")
}

func (te TurboEngine) Stall(){
	fmt.Println("Turbo Engine Stop...")
}