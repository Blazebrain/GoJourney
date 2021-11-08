package main

import (
	"fmt"
	"time"
)



func main (){
	//concurrency is the ability of an application to 
	//run multiple things at the same time

	go sayHello()
	time.Sleep(100 *time.Millisecond)
}




func sayHello (){
	fmt.Println("Hello")
}