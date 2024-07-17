package dto

import (
	"time"
)

type TodoRequestBody struct {
	Name string    `json:"name"`
	Done bool      `json:"done"`
	Time time.Time `json:"time"`
}

type TodoRequestParam struct {
	ID int `json:"id"`
}
