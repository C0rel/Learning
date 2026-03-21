package main

import "fmt"

type Book struct {
	Title    string
	Author   Author
	Pages    int
	Category string
}

type Author struct {
	FirstName string
	LastName  string
}

func NewBook(title, authorFirstName, authorLastName string, pages int, category string) Book {
	return Book{
		Title: title,
		Author: Author{
			FirstName: authorFirstName,
			LastName:  authorLastName,
		},
		Pages:    pages,
		Category: category,
	}
}
func main() {
	books := []Book{}
	books = append(books, NewBook("Преступление и наказание", "Фёдор Михайлович", "Достоевский", 627, "Роман"))
	fmt.Println(books)
}
