package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Rectangle struct{
	a,b float64
}

func (q Rectangle)Area() float64{
	S:=q.a*q.b
	return S
}

type Circle struct{
	R float64
}

func (q Circle)Area() float64{
	S:=math.Pow(q.R,2)*math.Pi
	return S
}
func main() {
    shapes:=[]Shape{
		Circle{R:2.0},
		Rectangle{a:2.0,b:3.0},
	}
	for _,sh:=range shapes{
		fmt.Println(sh.Area())
	}
}