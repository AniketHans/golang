package main

/*

// ****************** WITHOUT INTERFACES ********************
type englishBot struct{}

type spanishBot struct{}

// Here, in the below function definition, we are not using the reciever eb in the function so we can omit that also
// func (englishBot) getGreetings() string {} will also work
func (eb englishBot) getGreetings() string {
	// Custom logic for generating an English Greeting.
	return "Hello"
}

func (sb spanishBot) getGreetings() string {
	// Custom logic for generating a Spanish Greeting.
	return "Hola"
}

func printGreeting(eb englishBot){
	// this function has a similar logic to the one with sb object
	fmt.Println(eb.getGreetings())
}

// Go lang does not support function overloading, means we cannot have 2 functions with same name even if the agruments are different for them.
func printGreeting(sb spanishBot){
	// this function has a similar logic to the one with eb object
	fmt.Println(sb.getGreetings())
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

// *********************************************************

*/