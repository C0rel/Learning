package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("task.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return 
	}
	defer file.Close() 

	str := []byte("Данные")
	_, err = file.Write(str)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	fmt.Println("Запись выполнена успешно")
}