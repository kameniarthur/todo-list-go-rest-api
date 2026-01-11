package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

type User struct {
	ID   	 string `json:"id"`
	username string `json:"username"`
}

type Task struct {
	ID	 string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var (
	user = make(map[string]User)
	tasks = make(map[string]Task)
	token = make(map[string]string)

	usersMutex = &sync.Mutex{}
	tasksMutex = &sync.Mutex{}
	tokenMutex = &sync.Mutex{}
)

func main() {
	testUserID := uuid.New().String()
	user[testUserID] = User{ID: testUserID, Username: "testuser", Password: "password"}
	r := mux.NewRouter()

	log.Println("serveur demarre aur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))