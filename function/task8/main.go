package main

import "fmt"

func main() {
	arr := []int{4, 3, 6, 7, 1, 2, 8, 9, 5}
	fmt.Println("По возрастанию")
	sortInts(arr, less)
	fmt.Println("По убыванию")
	sortInts(arr, more)
}

func less(a, b int) bool {
	return a < b
}

func more(a, b int) bool {
	return a > b
}

func sortInts(slice []int, less func(a, b int) bool) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if !less(slice[j], slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
}