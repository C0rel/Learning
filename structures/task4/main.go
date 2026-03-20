package main

import "fmt"

type book struct {
	title  string
	author string
	pages  int
}

func findBookByAuthor(books []book, author string) []book{
	booksAutor:=[]book{}
	for i:=0;i<len(books);i++{
		if books[i].author == author{
			booksAutor = append(booksAutor,books[i])
		}
	}
	return booksAutor
}

func main() {
	books := []book{}

	books = append(books, book{"Преступление и наказание", "Фёдор Михайлович Достоевский", 672})
	books = append(books, book{"Сто лет одиночества", "Габриэль Гарсиа Маркес", 416})
	books = append(books, book{"Братья Карамазовы", "Фёдор Михайлович Достоевский", 800})
	
	fmt.Println(findBookByAuthor(books,"Фёдор Михайлович Достоевский"))
}
