package main

import "fmt"

type st string;

func main() {
	var cards []uint64 = []uint64{1, 2, 3}
	fmt.Println(cards)

	ret1,ret2:= func1("lol",5)
	fmt.Println(ret1,ret2)

	s1 := "a"
	s2 := st("b")
	fmt.Printf("%T",s1) // string
	fmt.Println()
	fmt.Printf("%T",s2) //main.st
}

func func1(s string,i int)(string,int){
	return s+"hi",i+10
}