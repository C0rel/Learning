package main

import (
	"errors"
	"fmt"
	"os"
)

func ReadConfig(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось загрузить конфиг: %w", err)
	}
	return nil
}

func main() {
	err := ReadConfig("файл_которого_нет.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Файл не найден:", err)
	} else if err != nil {
		fmt.Println("Другая ошибка:", err)
	} else {
		fmt.Println("Конфиг загружен успешно")
	}
}