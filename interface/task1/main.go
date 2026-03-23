package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct {
}

type Dog struct {
}

func (sp Cat) Speak() string {
	return "Мяу - Мяу"
}
func (sp Dog) Speak() string {
	return "Гав - Гав"
}
func main() {
	Animals := []Animal{Cat{},Dog{}}
	for _, animals := range Animals {
		voice:=animals.Speak()
		fmt.Println(voice)
	}
}