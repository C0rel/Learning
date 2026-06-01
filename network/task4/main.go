package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add",func(w http.ResponseWriter, r *http.Request){
		//http://localhost:8080/add?a=5&b=3
		a:=r.URL.Query().Get("a")
		b:=r.URL.Query().Get("b")
		if a==""||b==""{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error":"missing parameters"})
			return
		}
		aInt,err:=strconv.Atoi(a)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		bInt,err:=strconv.Atoi(b)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		result := map[string]int{"result": aInt+bInt}
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/sub",func(w http.ResponseWriter, r *http.Request){
		//http://localhost:8080/sub?a=5&b=3
		a:=r.URL.Query().Get("a")
		b:=r.URL.Query().Get("b")
		if a==""||b==""{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error":"missing parameters"})
			return
		}
		aInt,err:=strconv.Atoi(a)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		bInt,err:=strconv.Atoi(b)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		result := map[string]int{"result": aInt-bInt}
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/mul",func(w http.ResponseWriter, r *http.Request){
		//http://localhost:8080/mul?a=5&b=3
		a:=r.URL.Query().Get("a")
		b:=r.URL.Query().Get("b")
		if a==""||b==""{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error":"missing parameters"})
			return
		}
		aInt,err:=strconv.Atoi(a)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		bInt,err:=strconv.Atoi(b)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		result := map[string]int{"result": aInt*bInt}
		json.NewEncoder(w).Encode(result)
	})
	http.HandleFunc("/div",func(w http.ResponseWriter, r *http.Request){
		//http://localhost:8080/div?a=5&b=3
		a:=r.URL.Query().Get("a")
		b:=r.URL.Query().Get("b")
		if a==""||b==""{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error":"missing parameters"})
			return
		}
		aInt,err:=strconv.Atoi(a)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		bInt,err:=strconv.Atoi(b)
		 if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "invalid numbers"})
            return
        }
		if bInt==0{
			w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "division by zero"})
            return
		}
		result := map[string]float64{"result": float64(aInt)/float64(bInt)}
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080",nil)
}