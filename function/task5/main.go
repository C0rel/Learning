package main

import "fmt"

func main() {
	i:=nextInt()

	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	
}
func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}