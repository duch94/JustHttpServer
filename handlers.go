package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/duch94/JustHttpServer/clients"
)

type Handlers struct {
	Client *clients.MongoClient
}

func NewHandlers(mongoHost string, mongoPort string) (*Handlers) {
	// подключиться к бд, получить клиент 
	return nil
}

// RootHandler is used for authorization
func (h Handlers) RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is simple http server")
	fmt.Println(r)
}

// CreateUserHandler is handler for /createUser method
func (h Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorMessage := "CreateUser request was performed not by POST method"
		fmt.Fprintf(w, errorMessage)
		panic(fmt.Sprintln(errorMessage))
	}

	r.ParseForm()

	newUser := make(map[string]interface{})
	newUser["name"] = r.PostForm.Get("name")
	newUser["password"] = r.PostForm.Get("password")
	newUser["dob"] = r.PostForm.Get("dob")
	newUser["login"] = r.PostForm.Get("login")

	// надо возвращать http коды
	userID, err := h.Client.SendDocument("Main", "Users", newUser)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Created user %s\n", userID)
}

// ReadUserHandler is handler for /readUser method
func (h Handlers) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorMessage := "ReadUser request was performed not by GET method"
		fmt.Fprintf(w, errorMessage)
		panic(fmt.Sprintln(errorMessage))
	}

	url := r.URL.Query()
	login := url.Get("login")

	user, err := clients.GetDocumentByLogin("Main", "Users", login)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Read user: \n%s\n", user)
}

// UpdateUserHandler is handler for /updateUser method
func (h Handlers) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errorMessage := "UpdateUser request was performed not by PUT method"
		fmt.Fprintf(w, errorMessage)
		panic(fmt.Sprintln(errorMessage))
	}

	fmt.Fprintf(w, "Updated user")
}

// DeleteUserHandler is handler for /DeleteUser method
func (h Handlers) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		errorMessage := "DeleteUser request was performed not by DELETE method"
		fmt.Fprintf(w, errorMessage)
		panic(fmt.Sprintln(errorMessage))
	}

	fmt.Fprintf(w, "Deleted user")
}

// UserListHandler is handler for /userList method
func (h Handlers) UserListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorMessage := "UserList request was performed not by GET method"
		fmt.Fprintf(w, errorMessage)
		panic(fmt.Sprintln(errorMessage))
	}
	
	fmt.Fprintf(w, "Server is healthy!")
}

// HealthCheckHandler is handler for /healthCheck method
func (h Handlers) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is healthy!")
}
