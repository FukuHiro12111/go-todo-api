package model

import (
	"time"
)

type Task struct {
	ID       uint      `json:"id"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	Title    string    `json:"title"`
}
