package main

import (
	"api-todolist/src/connection"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var task Task

var prefrixPath = "/api/todo-list"

func main() {
	route := mux.NewRouter().StrictSlash(true)

	migrate := flag.Bool("migrate", false, "Create database")
	flag.Parse()

	if *migrate {
		err := connection.Migrations()

		if err != nil {
			log.Fatal(err)
		}
	}

	route.HandleFunc(prefrixPath+"/task", GetAllTasksController).Methods("GET")
	route.HandleFunc(prefrixPath+"/task/{id}", GetTaskController).Methods("GET")
	route.HandleFunc(prefrixPath+"/task", CreateTaskController).Methods("POST")
	route.HandleFunc(prefrixPath+"/task/{id}", UpdateTaskController).Methods("PUT")
	route.HandleFunc(prefrixPath+"/task/{id}", DeleteTaskController).Methods("DELETE")

	route.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", route))
}
