package main

import "fmt"

func main() {
	chA := make(chan string)
	chB := make(chan string)

	go func() {
		chA <- "A"
	}()
	go func() {
		chB <- "B"
	}()
	select {
	case valChA := <-chA:
		fmt.Println(valChA)
	case valChB:=<-chB:
		fmt.Println(valChB)
	}
}