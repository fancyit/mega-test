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
	//main method
	router.HandleFunc("/get-phrase-hash", c.MainMethod).Methods("POST")
	// additional method with details
	router.HandleFunc("/get-detailed-hash", c.DetailedMethod).Methods("POST")
	//if someone tries to access throe web, will get a notification of usage
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./source/static/")))
	//in case of errors while starting up the server
	if err := http.ListenAndServe(":"+port, router); nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("Server listens on port: ", port)
	}
}
