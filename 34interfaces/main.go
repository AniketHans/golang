// package main

// import "fmt"

// /*

// // ****************** WITHOUT INTERFACES ********************
// type englishBot struct{}

// type spanishBot struct{}

// // Here, in the below function definition, we are not using the reciever eb in the function so we can omit that also
// // func (englishBot) getGreetings() string {} will also work
// func (eb englishBot) getGreetings() string {
// 	// Custom logic for generating an English Greeting.
// 	return "Hello"
// }

// func (sb spanishBot) getGreetings() string {
// 	// Custom logic for generating a Spanish Greeting.
// 	return "Hola"
// }

// func printGreeting(eb englishBot){
// 	// this function has a similar logic to the one with sb object
// 	fmt.Println(eb.getGreetings())
// }

// // Go lang does not support function overloading, means we cannot have 2 functions with same name even if the agruments are different for them.
// func printGreeting(sb spanishBot){
// 	// this function has a similar logic to the one with eb object
// 	fmt.Println(sb.getGreetings())
// }

// func main() {

// 	eb := englishBot{}
// 	sb := spanishBot{}
// 	printGreeting(eb)
// 	printGreeting(sb)

// }

// // *********************************************************

// */

// // ************************* WITH INTERFACE **************************

// type bot interface{
// 	getGreetings() string
// 	demographics() string
// }

// type englishBot struct{}

// type spanishBot struct{}

// func (eb englishBot) getGreetings() string {
// 	// Custom logic for generating an English Greeting.
// 	return "Hello"
// }

// func (eb englishBot) demographics() string{
// 	return "Almost all over the world"
// }

// func (sb spanishBot) getGreetings() string {
// 	// Custom logic for generating a Spanish Greeting.
// 	return "Hola"
// }

// func printGreeting(b bot){
// 	fmt.Println(b.getGreetings())
// }

// func main() {

// 	eb := englishBot{}
// 	sb := spanishBot{}
// 	printGreeting(eb)
// 	printGreeting(sb)

// 	/*
// 		Hello
// 		Hola
// 	*/

// }

// // ********************************************************************/

package main

import "fmt"

// *********** struct A ************
   type A struct{}

   func (a A) func_common() string{
      return "This belongs to A"
   }

   func (a A) func_A() string{
      return "This is also belongs to A"
   }
   // ***********************************

   // ************ struct B *************
   type B struct{}

   func (b B) func_common() string{
      return "This belongs to B"
   }

   // ************************************


   // ************* INTERFACE *************

   type AB interface{
      func_common() string
   }
   type CD interface{
	func_common() string
   }
   // ************************************

   func call(ab AB){
      fmt.Println("This is called by AB")
   }
   func calling(cd CD){
      fmt.Println("This is called by CD")
   }

   func main(){
      a := A{}
      b := B{}
      call(a) // This is called by AB
      call(b) // This is called by AB
	  calling(a)
	  calling(b)
   }