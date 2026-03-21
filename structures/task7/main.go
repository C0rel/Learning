package main

import "fmt"

type book struct {
	title  string
	author Author
	pages  int
	category string
}

type Author struct {
	FirstName string
	LastName  string
	BirthYear string
}

func (b *book) AddPages(n int){
	b.pages+=n;
}

func main() {
	books := []book{}
	catalog:=make(map[string][]book)
	books = append(books,
		book{
			title: "Преступление и наказание",
			author: Author{
				FirstName: "Фёдор Михайлович",
				LastName:  "Достоевский",
				BirthYear: "11 ноября 1821 год",
			},
			category:"Роман",
			pages: 672})
	books = append(books,
		book{
			title: "Сто лет одиночества",
			author: Author{
				FirstName: "Габриэль Гарсиа",
				LastName:  "Маркес",
				BirthYear: "6 марта 1927 года",
			},
			category:"Детектив",
			pages: 416})

			for _,b:=range books{
				catalog[b.category]=append(catalog[b.category], b)
			}
			fmt.Println(catalog)
}
