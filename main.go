package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	dbCredentials := struct {
		host string
		port string
	}{"localhost", "27017"}

	handlers, err := NewHandlers(dbCredentials.host, dbCredentials.port)
	if err != nil {
		panic(err)
	}

	defer handlers.Client.Disconnect()

	router.HandleFunc("/", handlers.RootHandler).Methods("GET")
	router.HandleFunc("/createUser", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/readUser", handlers.ReadUserHandler).Methods("GET")
	router.HandleFunc("/updateUser", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/deleteUser", handlers.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/healthCheck", handlers.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/userList", handlers.UserListHandler).Methods("GET")

	http.Handle("/", router)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
