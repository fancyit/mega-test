package main

import (
	c "./source/controllers"
	m "./source/middlewares"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func init() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
}
func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.Use(m.UrlCheker)
	router.HandleFunc("/get-phrase-hash", c.MainMethod).Methods("POST")
	if err := http.ListenAndServe(":"+port, router); nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("Server listens on port: ", port)
	}
}
