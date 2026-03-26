package main

import (
	"fmt"
)


 
func SafeAccess(slice []int, index int) (value int, err error){
	defer func(){
		if r := recover(); r != nil {
			err = fmt.Errorf("выход за пределы слайса: %v", r)
		}
	}()

		return slice[index],nil
	
}
func main() {
    sl:=[]int{1,2,3,4,5,6,7}
	res,err:=SafeAccess(sl,3)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Элемент слайса: %d",res)
	}
}