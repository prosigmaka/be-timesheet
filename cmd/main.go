package main

import (
	"be-timesheet/pkg/config"
	"be-timesheet/pkg/db"
	"be-timesheet/pkg/handler"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	config.GetConfig()
}

func main() {
	DB := db.InitDB()
	h := handler.New(DB)
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/timesheets", h.GetAllTimesheets).Methods(http.MethodGet)
	router.HandleFunc("/timesheets/{id}", h.GetTimesheet).Methods(http.MethodGet)
	router.HandleFunc("/timesheets", h.AddTimesheet).Methods(http.MethodPost)
	router.HandleFunc("/timesheets/{id}", h.UpdateTimesheet).Methods(http.MethodPut)
	// router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	// http.ListenAndServe(":4000", router)
	http.ListenAndServe(":4000", handlers.CORS(headers, methods, origins)(router))
}
