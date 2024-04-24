package main

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectId    int64     `json:"projectId"`
	AssignedToId int64     `json:"assignedToId"`
	CreatedAt    time.Time `json:"createdAt"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Project struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateTaskPayload struct {
	Name         string `json:"name"`
	ProjectId    int64  `json:"projectId"`
	AssignedToId int64  `json:"assignedTo"`
}

type CreateProjectPayload struct {
	Name string `json:"name"`
}
