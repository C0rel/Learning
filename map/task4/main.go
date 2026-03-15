package main

import "fmt"

func main() {
	log:=map[string][]int{
		"Анна":{5,4,5},
		"Петр":{3,4,4},
		"Мария":{5,5,4},
	}
	for k,v:=range log{
		sum,count:=0,0
		for i:=0;i<len(v);i++{
			sum+=v[i]
			count++
		}
		fmt.Printf("Имя учащегося: %s, его средний балл: %.2f\n",k,float64(sum)/float64(count))
	}


}