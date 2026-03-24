package main

import "fmt"

type A struct{
	a int
}

func PrintAnything(v interface{}){
	fmt.Println(v)
}
func main() {
	a:=A{a:1}
	sl:=[]int{1,2,3,4,5}
    PrintAnything("Привет")
	PrintAnything(5)
	PrintAnything(a)
	PrintAnything(sl)
}