package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nazhard/go-simple-crud/pkg/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handlers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	log.Println("API running nyaa~")
	http.ListenAndServe(":3000", r)
}