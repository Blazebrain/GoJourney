package main

import "fmt"


func main (){
	switch 6 {
	case 1, 5, 7: 
			fmt.Println("the numbers are odd")
	case 2, 4, 6, 8: 
			fmt.Println("the numbers are even")
	default:
			fmt.Println("neither or higher scope")
		 
	}
}


