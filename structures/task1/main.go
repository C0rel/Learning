package main

import "fmt"

type book struct {
	title  string
	author string
	pages  int
}
func main() {
	book1 := book{"Преступление и наказание", "Фёдор Михайлович Достоевский", 672}
	book2 := book{"Сто лет одиночества", "Габриэль Гарсиа Маркес", 416}
	book3 := book{"Гарри Поттер и философский камень", "Дж. К. Роулинг", 432}
	fmt.Printf("Название: %s;Автор: %s;Страниц: %d\n",book1.title,book1.author,book1.pages)
	fmt.Printf("Название: %s;Автор: %s;Страниц: %d\n",book2.title,book2.author,book2.pages)
	fmt.Printf("Название: %s;Автор: %s;Страниц: %d\n",book3.title,book3.author,book3.pages)
}
