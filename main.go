package main

import "fmt"



func main (){

	// Performing Type switchning in Go
	var i interface {} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("i is an int") //Code will break and this code won't print
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default: 
	fmt.Println("i is another type")
		 
	}
}


