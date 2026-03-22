package main

import (
	"fmt"
)

func main() {
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)
	}

	res, err = safeDivide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", res)
	}
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("паника: %v", r)
		}
	}()

	if b == 0 {
		panic("деление на ноль")
	}
	result = a / b
	return
}