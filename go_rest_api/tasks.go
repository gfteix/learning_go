package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var errorNameRequired = errors.New("name is required")
var errorProjectIdRequired = errors.New("project id is required")
var errorUserIdRequired = errors.New("user id is required")

type TaskService struct {
	store Store
}

func NewTaskService(s Store) *TaskService {
	return &TaskService{store: s}
}

func (s *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", WithJWTAuth(s.handleCreateTask, s.store)).Methods("POST")
	r.HandleFunc("/Tasks/{id}", WithJWTAuth(s.handleGetTask, s.store)) //Methods("GET")
}

func (s *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid Request Payload"})
		return
	}

	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	task, err = s.store.CreateTask(task)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}

	WriteJson(w, http.StatusCreated, task)
}

func (s *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "id is required"})
		return
	}

	task, err := s.store.GetTask(id)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusNotFound, ErrorResponse{Error: "Task Not Found"})
		return
	}

	WriteJson(w, http.StatusOK, task)
}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errorNameRequired
	}

	if task.ProjectId == 0 {
		return errorProjectIdRequired
	}

	if task.AssignedToId == 0 {
		return errorUserIdRequired
	}

	return nil
}
