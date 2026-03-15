package main

import "fmt"

func main() {
	fmt.Println("Введите имя")
	name:=""
	fmt.Scanln(&name)
	runes:=[]rune(name)
	char:= make(map[rune]int)
	
	for _,ch := range runes{
		char[ch]++
	}
	for ch,count:=range char{
		fmt.Printf("буква %q,повторений: %d\n",ch,count)
	}

}