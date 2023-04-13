package main

import (
	"fmt"
	"log"
	"net/http"
	"todos/api"
	"todos/db"

	"github.com/gorilla/mux"
)

func main() {
	mydb := db.EstablishConnection()
	if mydb == nil {
		fmt.Printf("Something is wrong in database connection...\n")
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/todos", api.GetAllTodos(mydb)).Methods("GET")
	r.HandleFunc("/todos/{id}", api.GetTodosWithID(mydb)).Methods("GET")
	r.HandleFunc("/todos", api.CreateTodos(mydb)).Methods("POST")
	r.HandleFunc("/todos/{id}", api.UpdateTodos(mydb)).Methods("PUT")
	r.HandleFunc("/todos/{id}", api.DeleteTodos(mydb)).Methods("DELETE")

	fmt.Printf("Connection established at localhost:9090 port!\n")
	log.Fatal(http.ListenAndServe(":9090", r))
	mydb.Close()
}
