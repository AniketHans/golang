package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	//for loop
	for d:=0;d<len(days);d++{
		fmt.Printf("%v day is %v\n",d+1,days[d])
	}

	fmt.Println("***********************")

	// another for loop
	for i :=range days{
		fmt.Printf("%v day is %v\n",i+1,days[i])
	}

	fmt.Println("#######################")
	// for each kind of loop
	for i, data :=range days{
		fmt.Printf("%v day is %v\n",i+1,data)
	}

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$")

	// for loop used as while loop
	value := 1
	for value<=10{
		fmt.Println(value)
		value++
	}

	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^")
	// we can use break, continiue and goto as control flow as well

	//using goto. We can jump to any label using goto
	val := 1
	point:
		fmt.Println("The value of val is",val)
		val++
	if val<=5{
		goto point
	}

	

}