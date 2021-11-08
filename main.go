package main

import (
	"fmt"
)



func main(){
	sum(1,2,3,4,5)
}

func sum(values ...int){
	fmt.Println(values)
	result :=0


	for i:=0; i < len(values); i++{
		println(values[i])
		result += values[i]
	}

	fmt.Println("The result is", result)
}