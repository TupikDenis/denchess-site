package main

// set GO111MODULE=on
// go run ./cmd/web/.

import (
	"log"
	"net/http"

	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/gorilla/mux"
)

func HandleFunc() {
	router := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	router.HandleFunc("/", home)
	router.HandleFunc("/nopage", pageNotFound)
	router.HandleFunc("/news", news)

	router.HandleFunc("/users", handlers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", handlers.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}

func ServerStart() {

}

func main() {
	HandleFunc()
}
