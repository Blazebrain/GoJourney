package main

import (
	"fmt"
)



func main (){
	for i:= 0; i<=5; i++{
		sayMessage("Hello Go", i)

	}

}
func sayMessage(msg string, i int){
	fmt.Println(msg)
	fmt.Println("The value of index is", i)
}

