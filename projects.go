package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequiredP error = errors.New("name is required")

type ProjectService struct {
	store Store
}

// Constructor
func NewProjectService(s Store) *ProjectService {
	return &ProjectService{store: s}
}

// Methods
func (ps *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", ps.handleCreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", ps.handleGetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", ps.handleDeleteProject).Methods("DELETE")
}

func (ps *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	if err := validateProjectPayload(project); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	p, err := ps.store.CreateProject(project)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating the project"})
		return
	}

	WriteJSON(w, http.StatusCreated, p)
}

func (ps *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {

}

func (ps *ProjectService) handleDeleteProject(w http.ResponseWriter, r *http.Request) {
	
}

func validateProjectPayload(p *Project) error {
	if p.Name == "" {
		return errNameRequiredP
	}

	return nil
}