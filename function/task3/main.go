package main

import (
	"fmt"
)

func main() {
	var (
		a, b float64
	)
	fmt.Println("Введите ширину и длину прямоугольника через пробел")
	fmt.Scanln(&a,&b)
	fmt.Printf("Площадь: %.2f", Area(a,b))

}

func Area(width,height float64)(res float64){
	res = width*height
	return 
}
