package main

import (
	c "./source/controllers"
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
	//router.Use(m.UrlCheker)
	//staticDir := "./static"
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./source/static/")))
	router.HandleFunc("/get-phrase-hash", c.MainMethod).Methods("POST")
	router.HandleFunc("/get-detailed-hash", c.DetailedMethod).Methods("POST")
	if err := http.ListenAndServe(":"+port, router); nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("Server listens on port: ", port)
	}
}
