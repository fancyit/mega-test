package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Pass struct {
  InputStr string `json:"password"`  
  HashedPass int
}

func main() {   
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
	var pass Pass
	json.NewDecoder(r.Body).Decode(&pass)
        fmt.Fprintf(w, "%s", pass.InputStr)
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8081", nil)
}


