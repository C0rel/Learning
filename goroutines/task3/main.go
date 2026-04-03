package main

import "fmt"

func main() {
	ch := make(chan string)
	go func ()  {
		ch<-"Привет из горутины"
	}()
		
	fmt.Println(<-ch)	
}