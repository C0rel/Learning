package main

import (
	"fmt"
)

func Step3()error{
	return fmt.Errorf("Ошибка в третьем шаге")
}
func Step2()error{
	err:=Step3()
	if err!=nil{
		return fmt.Errorf("Step2,%w",err)
	}
	return nil
}
func Step1()error{
	err:=Step2()
	if err!=nil{
		return fmt.Errorf("Step1,%w",err)
	}
	return nil
}
func main() {
	err:=Step1()
	if err!=nil{
		fmt.Printf("Цепочка ошибок: %s", err)
	}
}