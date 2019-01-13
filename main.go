package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/createUser", CreateUserHandler)
	router.HandleFunc("/readUser", ReadUserHandler)
	router.HandleFunc("/updateUser", UpdateUserHandler)
	router.HandleFunc("/deleteUser", DeleteUserHandler)
	router.HandleFunc("/healthCheck", HealthCheckHandler)

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}

