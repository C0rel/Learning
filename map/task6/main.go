package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	phoneBook:=map[string]string{}
	scanner := bufio.NewScanner(os.Stdin)
	var (
		name string
		number string
		choice int
		)
	for{
	fmt.Print("\033[2J\033[H")	
	fmt.Println("1.Добавить контакт")
	fmt.Println("2.Найти контакт")
	fmt.Println("3.Удалить контакт")
	fmt.Println("4.Вывести все контакты ")
	fmt.Println("5.Выход")
	fmt.Println("")
	fmt.Println("Выберите действие")
	fmt.Scanln(&choice)

		switch choice{
			case 1:
				fmt.Print("\033[2J\033[H")
				fmt.Println("Введите имя контакта")
				scanner.Scan()
				name=scanner.Text()
				fmt.Println("Введите номер контакта")
				scanner.Scan()
				phoneBook[name]=scanner.Text()
			 
			case 2:
				fmt.Print("\033[2J\033[H")
				fmt.Println("Введите имя контакта")
				scanner.Scan()
				name=scanner.Text()
				numSearch,ok:=phoneBook[name]
				if ok{
					fmt.Println("Найден контакт")
					fmt.Printf("%s : %s\n",name,numSearch)
				}else{
					fmt.Println("Контакт не найден")
				}
				fmt.Println("Нажмите Enter, чтобы продолжить...")
    			scanner.Scan()
			case 3:
				fmt.Print("\033[2J\033[H")
				fmt.Println("Введите имя контакта,который хотите удалить")
				scanner.Scan()
				name=scanner.Text()
				delete(phoneBook,name)
			case 4:
				fmt.Print("\033[2J\033[H")
				if len(phoneBook)==0{
					fmt.Println("Телефонная книга пуста")
				}else{
					for name,number = range phoneBook {
						fmt.Printf("%s : %s\n",name,number)
					} 
					}
				fmt.Println("Нажмите Enter, чтобы продолжить...")
    			scanner.Scan()
			case 5:
				os.Exit(0)
	}
}
	
}