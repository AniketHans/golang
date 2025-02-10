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

	func2("a","b")
	func3(1,"a","b","c",1.2)
}

func func1(s string,i int)(string,int){
	return s+"hi",i+10
}

func func2(s,p string){
	fmt.Printf("\n%T,%T",s,p)
}

func func3(a int, b,c,d string,e float32){
	fmt.Printf("\n%T,%T,%T,%T,%T",a,b,c,d,e)
}