package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
	ch := make(chan int)
	wg.Add(2)
	//Receiving goroutine
	//the parameter makes it a receive only goroutine
	//so data can't be sent from this goroutine
	go func (ch <-chan int){
		i:= <- ch
		fmt.Println(i)
	
		wg.Done()
	}(ch)
	//sending goroutine
	//the parameter makes it a send only channel goroutine, 
	//so data can't be received in this goroutine
	go func (ch chan <- int){
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}