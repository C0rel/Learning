package main

import "fmt"

type Reader interface{
	Read() string
}
type Writer interface{
	Write(s string)
}
type ReadWriter interface{
	Reader;Writer
}

type File struct{
}

func (f File)Read() string{
	return "Данные из файла"
}

func (f File)Write(s string){
	fmt.Println(s)
}

func main() {
   rw:=[]ReadWriter{File{}}
   for _,v:=range rw{
	fmt.Println(v.Read())
	v.Write("Данные")
   }
}