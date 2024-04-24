package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectService struct {
	store Store
}

func NewProjectService(s Store) *ProjectService {
	return &ProjectService{store: s}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", WithJWTAuth(s.handleCreateProject, s.store)).Methods("POST")
	r.HandleFunc("/projects/{id}", WithJWTAuth(s.handleGetProject, s.store)) //Methods("GET")
}

func (s *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid Request Payload"})
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON"})
		return
	}

	if project.Name == "" {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	project, err = s.store.CreateProject(project)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating project"})
		return
	}

	WriteJson(w, http.StatusCreated, project)
}

func (s *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "id is required"})
		return
	}

	project, err := s.store.GetProject(id)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusNotFound, ErrorResponse{Error: "Project Not Found"})
		return
	}

	WriteJson(w, http.StatusOK, project)
}
