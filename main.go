package main

import "fmt"


func main(){

	number := 50
	guess := 70

	if guess < number {
		fmt.Println("Too low")
	} else  if guess > number {
		fmt.Println("Too high")
	} else{ 
		fmt.Println("Perfect!")
	}

	fmt.Println(number<=guess, number>=guess, number !=guess)

}

func returnTrue() bool{
	fmt.Println("returning true")
	return true
}



