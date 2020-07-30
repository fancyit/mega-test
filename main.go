package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	c "./source/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	//load .env on startup
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}
func main() {

	var port string
	// read .env for port
	envPort, exists := os.LookupEnv("PORT")
	if exists {
		port = envPort
	} else {
		port = "900"
	}

	router := mux.NewRouter()

	//main method
	router.HandleFunc("/get-phrase-hash", c.MainMethod).Methods("POST")

	// additional method with details
	router.HandleFunc("/get-detailed-hash", c.DetailedMethod).Methods("POST")

	//if someone tries to access throe web, will get a notification of usage
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./source/static/")))

	log.Println("Server listens on port: ", port)

	//in case of errors while starting up the server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
