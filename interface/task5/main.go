package main

import (
	"fmt"
)

func Describe(i interface{}){
	switch value:=i.(type){
	case string:
		fmt.Println(len(value))
	case int:
		fmt.Println(value*value)
	default:
		fmt.Println("Неизвестный тип")
	}
	
}
func main() {
	sl:=[]int{1,2,3,4,5}
    Describe("Hello")
	Describe(5)
	Describe(sl)
}