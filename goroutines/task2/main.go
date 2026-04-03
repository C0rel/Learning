package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	num:= func(){
		for i:=0;i<3;i++{
			fmt.Println(i+1)
		}
		wg.Done()
	}
	go num()
	wg.Wait()	
}