package main

import "fmt"

func main() {
	double := multiplier(2)
	fmt.Println(double(5))
}

func multiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}