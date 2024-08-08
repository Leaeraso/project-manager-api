package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name is required")
var errProjectIDRequired = errors.New("project id is required")
var errUserIDRequired = errors.New("user id is required")

type TasksService struct {
	store Store
}

// Constructor
func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

// Methods
func (ts *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", WithJWTAuth(ts.handleCreateTask, ts.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", WithJWTAuth(ts.handleGetTask, ts.store)).Methods("GET")
}

func (ts *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading the request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	t, err := ts.store.CreateTask(task)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "error creating the task"})
		return
	}

	WriteJSON(w, http.StatusCreated, t)
}

func (ts *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	// recuperamos las variables de la http request
	vars := mux.Vars(r)
	// almacenamos la variable id
	id := vars["id"]

	t, err := ts.store.GetTask(id)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "task not found"})
		return
	}

	WriteJSON(w, http.StatusOK, t)
}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil
}
