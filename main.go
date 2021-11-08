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
	go func (){
		i:= <- ch
		fmt.Println(i)
		wg.Done()
	}()
	//sending goroutine
	go func (){
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}