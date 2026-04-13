package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go func() {
		ch <- 100
		wg.Done()
	}()
	go func() {
		ch <- 200
		wg.Done()
	}()
	go func(){
		wg.Wait()
		close(ch)
	}()
	sum:=0
	for value:=range ch{
		sum+=value
		fmt.Printf("Значение: %d,\n",value)
	}
	fmt.Printf("Сумма: %d.\n",sum)
}