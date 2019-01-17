package main

import (
	"fmt"
	"net/http"
	"context"
	"time"
	"github.com/duch94/JustHttpServer/clients"
)

// Handlers is object with request handler methods
type Handlers struct {
	Client *clients.MongoClient
}

// NewHandlers is constructor of Handlers object
func NewHandlers(mongoHost string, mongoPort string) (*Handlers, error) {
	// подключиться к бд, получить клиент 
	var (
		h Handlers
		err error
	)
	h.Client, err = clients.NewMongoClient(mongoHost, mongoPort)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// RootHandler is used for authorization
func (h *Handlers) RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is simple http server")
	fmt.Println(r)
}

// CreateUserHandler is handler for /createUser method
func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	newUser := make(map[string]interface{})
	newUser["name"] = r.PostForm.Get("name")
	newUser["password"] = r.PostForm.Get("password")
	newUser["dob"] = r.PostForm.Get("dob")
	newUser["login"] = r.PostForm.Get("login")

	// надо возвращать http коды
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userID, err := h.Client.SendDocument(ctx, "Main", "Users", newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated) 
	fmt.Fprintf(w, "Created user %s\n", userID)
}

// ReadUserHandler is handler for /readUser method
func (h *Handlers) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	login := url.Get("login")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user, err := h.Client.GetDocumentByLogin(ctx, "Main", "Users", login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Read user: \n%s\n", user)
}

// UpdateUserHandler is handler for /updateUser method
func (h *Handlers) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	login := url.Get("login")
	updatedKey := url.Get("updatedKey")
	updatedValue := url.Get("updatedValue")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	updatedNumber, err := h.Client.UpdateDocumentByLogin(ctx, "Main", "Users", login, 
												updatedKey, updatedValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Updated %d user documents\n", updatedNumber)
}

// DeleteUserHandler is handler for /DeleteUser method
func (h *Handlers) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	login := url.Get("login")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	deletedUsersNum, err := h.Client.DeleteDocumentByLogin(ctx, "Main", "Users", login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Deleted %d users\n", deletedUsersNum)
}

// UserListHandler is handler for /userList method
func (h *Handlers) UserListHandler(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	users, err := h.Client.GetAllDocuments(ctx, "Main", "Users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	fmt.Fprintf(w, "User list:\n%s\n", users)
}

// HealthCheckHandler is handler for /healthCheck method
func (h *Handlers) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is healthy!")
}
