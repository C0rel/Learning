package main

import (
	"errors"
	"fmt"
)

func main() {
    var(
        a,b,op int
    )
    fmt.Println("Введите два значения через пробел")
    fmt.Scanf("%d %d",&a,&b)
	fmt.Scanln()
    fmt.Println("Выберите операцию")
    fmt.Println("1 - Сложение")
    fmt.Println("2 - Вычитание")
    fmt.Println("3 - Умножение")
    fmt.Println("4 - Деление")
    fmt.Scanf("%d",&op)
    switch op {
    case 1:
        fmt.Println(Add(a,b))
         
    case 2:
        fmt.Println(Subtract(a,b))
       
    case 3:
        fmt.Println(Multiply(a,b))
         
    case 4:
        res, err := Divide(a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
        
    default:
        fmt.Println("Неизвестное значение")
    }
    
}

func Add(a int,b int)int{
    return a+b
}

func Subtract(a int,b int)int{
    return a-b
}

func Multiply(a int,b int)int{
    return a*b
}

func Divide(a int,b int)(int,error){
    if b!=0{
    return a/b , nil;
    }else{
        return 0 ,errors.New("Деление на ноль")
    }
}