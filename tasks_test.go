package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Vamos a testear los endpoints y ver si funcionan antes de pasarlos a produccion
func TestCreateTask(t *testing.T) {
	//MemoryStorage
	ms := &MockStore{}

	service := NewTasksService(ms)

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := &Task{
			Name: "",
		}
		//pasamos payload a un []byte
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		//creamos el request
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		//creamos el response recorder
		rr := httptest.NewRecorder()
		r := mux.NewRouter()
		//manejamos la ruta
		r.HandleFunc("/tasks", service.handleCreateTask)
		// Inicializamos
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code, it should fail!")
		}
	})

	t.Run("should create a task", func(t *testing.T) {
		payload := &Task{
			Name:         "Creating a REST API in Go",
			ProjectID:    1,
			AssignedToID: 44,
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()

		r.HandleFunc("/tasks", service.handleCreateTask)

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	// TODO: all the tests
}

func TestGetTask(t *testing.T) {
	ms := &MockStore{}

	service := NewTasksService(ms)

	t.Run("should return a task", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/tasks/44", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()

		r.HandleFunc("/tasks/{id}", service.handleGetTask)

		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error("invalid status code, it shouldn't fail")
		}
	})
}
