package main

import (
	"fmt"
	"net/http"
	"github.com/duch94/JustHttpServer/clients"
)

// RootHandler is used for authorization
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is simple http server")
	fmt.Println(r)
}

// CreateUserHandler is handler for /createUser method
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	newUser := make(map[string]interface{})
	newUser["name"] = "name"
	newUser["password"] = "password"
	newUser["dob"] = "01-01-1970"
	newUser["login"] = "email@mailbox.com"

	// надо возвращать http коды
	userID, err := clients.SendDocument("Main", "Users", newUser)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Created user %s", userID)
}

// ReadUserHandler is handler for /readUser method
func ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Read user")
}

// UpdateUserHandler is handler for /updateUser method
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updated user")
}

// DeleteUserHandler is handler for /DeleteUser method
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted user")
}

// UserListHandler is handler for /userList method
func UserListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is healthy!")
}

// HealthCheckHandler is handler for /healthCheck method
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is healthy!")
}
