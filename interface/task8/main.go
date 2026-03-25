package main

import "fmt"

type Counter interface{
	Increment()
}

type CounterImpl struct{
	value int
}

func (v *CounterImpl)Increment(){
	v.value++
	fmt.Println(v.value)
}
func main() {
    res:=[]Counter{
		&CounterImpl{value:1},
	}
	for _,v:=range res{
		v.Increment()
	}
}