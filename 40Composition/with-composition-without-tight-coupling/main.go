package main

func main() {
	c1 := Car{TurboEngine{}, Transmission{}, Steering{}, "Ferrari"}
	c2 := Car{Engine{}, Transmission{}, Steering{}, "Pagani"}
	t1 := Truck{Engine{}, AdvancedTransmission{}, Steering{}, "Tata"}
	t2 := Truck{Engine{}, Transmission{}, Steering{}, "Lamborgini"}

	c1.Start()
	c2.Start()

	t1.ShiftUp()
	t2.ShiftUp()
}