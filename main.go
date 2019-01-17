package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	dbCredentials := struct {
		host string
		port string
	}{ "localhost", "27017" }
	
	h, err := NewHandlers(dbCredentials.host, dbCredentials.port)
	if err != nil {
		panic(err)
	}
	defer h.Client.Disconnect()

	router.HandleFunc("/", h.RootHandler).Methods("GET")
	router.HandleFunc("/createUser", h.CreateUserHandler).Methods("POST")
	router.HandleFunc("/readUser", h.ReadUserHandler).Methods("GET")
	router.HandleFunc("/updateUser", h.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/deleteUser", h.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/healthCheck", h.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/userList", h.UserListHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}

