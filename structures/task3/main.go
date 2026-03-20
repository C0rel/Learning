package main

import "fmt"

type book struct {
	title  string
	author string
	pages  int
}

func (b book) IsThicker(other book) bool {
	return b.pages > other.pages
}

func main() {
	books := []book{}

	books = append(books, book{"Преступление и наказание", "Фёдор Михайлович Достоевский", 672})
	books = append(books, book{"Сто лет одиночества", "Габриэль Гарсиа Маркес", 416})
	result := books[0].IsThicker(books[1])
	if result {
		fmt.Println("Книга 1 тольще")
	}else{
		fmt.Println("Книга 2 тольще")
	}

}
