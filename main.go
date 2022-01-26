package main

import (
	"log"
	"net/http"
	"timesheet-app/pkg/db"
	"timesheet-app/pkg/handlers"

	"github.com/gorilla/mux"
)



func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/projects",h.GetAllProject ).Methods(http.MethodGet)
	router.HandleFunc("/projects/{id}",h.GetProject ).Methods(http.MethodGet)
	router.HandleFunc("/projects",h.AddProject).Methods(http.MethodPost)
	router.HandleFunc("/projects/{id}", h.UpdateProject).Methods(http.MethodPut)
	router.HandleFunc("/projects/{id}", h.DeletProject).Methods(http.MethodDelete)


	log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}