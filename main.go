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

	router.HandleFunc("/", h.RootHandler)
	router.HandleFunc("/createUser", h.CreateUserHandler)
	router.HandleFunc("/readUser", h.ReadUserHandler)
	router.HandleFunc("/updateUser", h.UpdateUserHandler)
	router.HandleFunc("/deleteUser", h.DeleteUserHandler)
	router.HandleFunc("/healthCheck", h.HealthCheckHandler)
	router.HandleFunc("/userList", h.UserListHandler)

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}

