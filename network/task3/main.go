package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
type Value struct{
	A int
	B int
}

func main() {
	go func(){
	http.HandleFunc("/sum",func(w http.ResponseWriter, r *http.Request) {
		var data map[string]int
		err := json.NewDecoder(r.Body).Decode(&data)
       	if err != nil {
            
            http.Error(w, "Невалидный JSON", http.StatusBadRequest)
            return
        }
		a:=data["A"]
		b:=data["B"]
		result:=a+b
		answer := map[string]int{"result": result}
		json.NewEncoder(w).Encode(answer)
		defer r.Body.Close()
	})
	http.ListenAndServe(":8081",nil)
	}()
 	time.Sleep(3 * time.Second)	
	
	val:=Value{A:5,B:4}
	jsonVal,err:=json.Marshal(val)
	 if err != nil {
        fmt.Println("Ошибка создания JSON:", err)
        return
    }
	resp,err:=http.Post(
		"http://localhost:8081/sum",
		"application/json",
        bytes.NewBuffer(jsonVal),
	)
	 if err != nil {
        fmt.Println("Ошибка отправки:", err)
        return
    }
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Scanln() 
}