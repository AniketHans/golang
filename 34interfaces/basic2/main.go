package main

// **************** Complicated Eg ********************

// import "fmt"

// type A struct{}
// type B struct{}

// func (a A) func_common() string {
// 	return "Called by A"
// }

// func (a B) func_common() string {
// 	return "Called by B"
// }

// type AB interface {
// 	func_common() string
// }

// type C struct {
// 	c  string
// 	ab AB
// }

// func main() {
// 	a := A{}
// 	c := C{c: "Hi", ab: a}
// 	fmt.Println(c) // {Hi {}}
// }

// *******************************************

// *************** MORE COMPLICATED E.g. **********************
import "fmt"

type A struct{}
type B struct{}

func (a A) func_common_1() string {
	return "Called by A"
}

func (b B) func_common_1() string {
	return "Called by B"
}


func (a A) func_common_2() string {
	return "Called by A 2"
}


type Int1 interface{
	func_common_1() string
}

type Int2 interface{
	func_common_2() string
}

/* 
	Here, a struct will be called a member of Int12 only and only if it is a member of both Int1 and Int2.
	For a struct to be member of Int1, it should have an associated function with reference `func_common_1() string`
	For a struct to be member of Int2, it should have an associated function with reference `func_common_2() string`
	Thus it means, for a struct to be member of Int12, it should have both `func_common_1() string` and `func_common_2() string` associated with them.
	Here, struct B only has `func_common_1() string` associated with it.
	But, struct A has both so any struct A is the member of Int12 interface.
	ProStruct need a member of Int12 thus object of struct A can be used here.
*/
type Int12 interface {
	Int1
	Int2
}

type ProStruct struct {
	str  string
	inter Int12
}

func main() {
	a := A{}
	ps := ProStruct{str:"Holaa", inter: a}
	fmt.Println(ps) // {Holaa {}}
}

// *******************************************