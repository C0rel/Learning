package main

import (
	"fmt"
	"sync"
)
var mutex sync.Mutex  

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	mutex.Lock()
	c.value++
	mutex.Unlock()
}

func (c *Counter) Value() int {
	mutex.Lock()
	defer mutex.Unlock()
	return c.value
}
func main() {
	var wg sync.WaitGroup
	copyStuct := Counter{0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done() 
			copyStuct.Increment()
			}()
		
	}
	  wg.Wait()
	  fmt.Println(copyStuct.Value())
}