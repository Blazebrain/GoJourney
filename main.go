package main

import "fmt"



func main(){
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World!"))
}

//Interface describes behaviours not types
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

//creating methods: basically functions with a context

func (cw ConsoleWriter) Write(data []byte) (int, error){
	n,err := fmt.Println(string(data))
	return n,err
}