package main

import (
	"fmt"
	"math"
)

type NegativeNumberError struct{
	num float64
}
func (n NegativeNumberError)Error()string{
	err:=fmt.Sprintf("Отрицательное число:%.2f",n.num)
	return err 
}
func SqrtPositive(x float64)(float64,error){
	if x<0{
		return 0,NegativeNumberError{num: x}
	}else{
		return math.Sqrt(x),nil
	}
}
func main() {
	v,err:=SqrtPositive(-4)
	if err!=nil{
		fmt.Printf("Ошибка: %s",err)
	}else{
		fmt.Printf("Результат: %.2f",v)
	}

}