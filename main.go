package main

import "fmt"



func main (){
	s := []int{1,2,3 ,4}

	for k, v := range s{
		fmt.Println(k, v)
	}
}


