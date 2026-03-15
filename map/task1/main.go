package main

import "fmt"

func main() {
	people:=map[string]int{
		"Анна":25,
		"Петр": 30, 
		"Мария": 28,
	}
	name:=""
	fmt.Printf("%v\n",people)
	fmt.Print("Введите имя:")
	fmt.Scanln(&name)
	age,exist:=people[name]
	if exist{
		fmt.Printf("Резултат найден.\nИмя: %s, возраст: %d",name,age)
	}else{
		fmt.Println("Ничего не найдено")
	}
}
