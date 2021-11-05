package main

import "fmt"



func main (){
	//Defer : invoke function but defer it's applicationi to later
		fmt.Println("start")
		defer fmt.Println("middle")
		fmt.Println("end")

}	


