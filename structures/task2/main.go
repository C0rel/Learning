package main

import (
	"fmt"
)

type book struct {
	title  string
	author string
	pages  int
}

func main() {
	books := []book{}

	books = append(books,book{"Преступление и наказание", "Фёдор Михайлович Достоевский", 672})
	books = append(books,book{"Сто лет одиночества", "Габриэль Гарсиа Маркес", 416})
	books = append(books,book{"Гарри Поттер и философский камень", "Дж. К. Роулинг", 432})
	printBooks(books)
}

func printBooks(books []book) {
	 for _, b := range books{
	fmt.Printf("Название: %s;Автор: %s;Страниц: %d\n", b.title, b.author, b.pages)
	}
}
