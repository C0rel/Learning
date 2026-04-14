package main

import (
	"fmt"
	"sync"
)

func worker1(jobs <-chan int,result chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range jobs {
		result <- task * task
	}

}
func worker2(jobs <-chan int,result chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range jobs {
		result <- task * task
	}

}

func worker3(jobs <-chan int,result chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range jobs {
		result <- task * task
	}

}




func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	result := make(chan int,10)
	wg.Add(3)
	go worker1(ch,result,&wg)
	go worker2(ch,result,&wg)
	go worker3(ch,result,&wg)
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
	close(ch)
	wg.Wait()
	close(result)

	for res := range result{
		fmt.Println(res)
	}
}