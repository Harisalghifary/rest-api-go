package main

import (
	"log"
	"net/http"

	"github.com/Harisalghifary/rest-api-go/apps"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", apps.HomeLink)
	router.HandleFunc("/users", apps.CreateUser).Methods("POST")
	router.HandleFunc("/users", apps.GetAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", apps.GetOneUser).Methods("GET")
	router.HandleFunc("/users/{id}", apps.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", apps.DeleteOneUser).Methods("DELETE")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
