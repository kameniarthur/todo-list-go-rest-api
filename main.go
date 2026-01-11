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
}

func createTaskHandler(w http.ResposeWriter, r *http.Request) {	
	var newTask Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newTask.ID = uuid.New().string()
	tasksMutex.Lock()
	tasks[newTask.ID] = newTask
	tasksMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasksMutex.Rlock()
	defer tasksMutex.RUnlock()

	var tasksList []Task
	for _, task := range tasks {
		tasksList = append(tasksList, task)
	}
	tasksMutex.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasksList)
}

func taskDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	task, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	swich r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(task)
		case http.MethodPut:
			var updatedTask Task
			if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}
			updatedTask.ID = taskID
			updatedTask.ownerID = task.ownerID
			tasks[taskID] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
		case http.MethodDelete:
			delete(tasks, taskID)
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}