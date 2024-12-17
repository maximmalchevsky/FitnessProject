package models

type Message struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Id struct {
	Id int `json:"id"`
}
