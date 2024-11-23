package main

import "fmt"

func adder(val1 int, val2 int) int {
	return (val1 + val2)
}
func proAdder(values ...int)(int,string){
	s:=0
	for _, val:=range values{
		s+=val
	}
	return s,"Yo!! What's up??"
}

func main() {
	fmt.Println(adder(3, 5))
	fmt.Println(proAdder(3,4,5,6,7,8))
}