package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1,2,3,4,5))
}
func sum(nums ...int)int{
	summa:=0
	for _,v:= range nums{
		summa+=v
	}
	return summa
}
