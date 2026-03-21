package main

import "fmt"

type book struct {
	title  string
	author Author
	pages  int
}

type Author struct {
	FirstName string
	LastName  string
	BirthYear string
}

func main() {
	books := []book{}

	books = append(books,
		book{
			title: "Преступление и наказание",
			author: Author{
				FirstName: "Фёдор Михайлович",
				LastName:  "Достоевский",
				BirthYear: "11 ноября 1821 год",
			},
			pages: 672})
	books = append(books,
		book{
			title: "Сто лет одиночества",
			author: Author{
				FirstName: "Габриэль Гарсиа",
				LastName:  "Маркес",
				BirthYear: "6 марта 1927 года",
			},
			pages: 416})
	fmt.Println(books)
}
