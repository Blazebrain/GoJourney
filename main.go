package main

import "fmt"


func main (){
	i:= 2+3; 
	switch  {
	case i<=10:  
			fmt.Println("the numbers are odd")
			fallthrough
	case i<=20: 
			fmt.Println("the numbers are even")
	default:
			fmt.Println("neither or higher scope")
		 
	}
}


