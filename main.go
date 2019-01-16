package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	
	h, err := NewHandlers("localhost", "27017")
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/", h.RootHandler)
	router.HandleFunc("/createUser", h.CreateUserHandler)
	router.HandleFunc("/readUser", h.ReadUserHandler)
	router.HandleFunc("/updateUser", h.UpdateUserHandler)
	router.HandleFunc("/deleteUser", h.DeleteUserHandler)
	router.HandleFunc("/healthCheck", h.HealthCheckHandler)

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}

