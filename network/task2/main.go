package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		name:=r.URL.Query().Get("name")
		if name==""{
			name = "Guest"
		}
		str := map[string]string{"message": fmt.Sprintf("Hello, %s", name)}
		json.NewEncoder(w).Encode(str)
	})

	http.ListenAndServe(":8080",nil)
}