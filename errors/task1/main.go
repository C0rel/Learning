package main

import (
	"errors"
	"fmt"
)

func Divide(a,b float64) (float64,error){
	if b!=0{
		return a/b , nil
	}else{ return 0,errors.New("Деление на ноль")}
}
func main() {
    v,err:=Divide(2,0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", v)
	}
}