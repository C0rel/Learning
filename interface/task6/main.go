package main

import (
	"fmt"
)

type Person struct{
	name string
	age int
}
func (v Person)String()string{
	return fmt.Sprintf("Имя: %s, Возраст: %d", v.name, v.age)
}
func main() {
    p := Person{name:"Павел",age:22}
	fmt.Println(p)
}