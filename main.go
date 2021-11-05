package main

import "fmt"



func main (){
	//Defer : invoke function but defer it's applicationi to later
		defer fmt.Println("start")
		defer fmt.Println("middle")
		defer fmt.Println("end")

		//LIFO order:  if there is deferred in all of them
		//Last In First Out

}	


