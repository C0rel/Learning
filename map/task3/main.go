package main

import "fmt"

func main() {
	dictionary := map[string]string{
		"красивый":   "прекрасный",
		"говорить":   "болтать",
		"большой":    "огромный",
		"быстро":     "стремительно",
		"интересный": "занимательный",
	}
	fmt.Println("Введите слово")
	word:=""
	fmt.Scanln(&word)
	synonym,res:=dictionary[word]
	if res{
		fmt.Println(synonym)
	}else{
		fmt.Println("Ничего не найдено")
	}

}