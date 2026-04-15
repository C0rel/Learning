package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	status := map[string]string{"status": "ok"}   
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		//w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	})
		
	http.ListenAndServe(":8080",nil)
}