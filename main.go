package main

import (
	"fmt"
)



func main (){
		var a int = 42
		// & refers to address of * refers to de-referencing operator
		var b *int = &a
		fmt.Println(a,*b)
}


