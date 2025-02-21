// course name:- course/go-the-complete-developers-guide/learn/lecture/38476714#overview
package main

import "fmt"

type ShapeArea interface{
	getArea() float64
}

type Triangle struct{
	base float64
	height float64
}
type Square struct{
	side float64
}

func (t Triangle) getArea() float64{
	return 0.5*t.base*t.height
}

func (s Square) getArea() float64{
	return s.side*s.side
}


func PrintArea(sa ShapeArea){
	fmt.Println(sa.getArea())
}

func main(){
	t := Triangle{
		base: 10,
		height: 10,
	}

	s := Square{
		side:10,
	}

	PrintArea(t) //50
	PrintArea(s) //100

}