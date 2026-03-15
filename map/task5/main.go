package main

import "fmt"

func main() {
	ages := map[string]int{
		"Анна":  25,
		"Петр":  30,
		"Мария": 28,
		"Иван":  25,
	}
	agesRevers := map[int][]string{}

	for name, age := range ages {
		agesRevers[age] = append(agesRevers[age], name)
	}
	for age, names := range agesRevers {
		fmt.Printf("%d: %v\n", age, names)
	}
}