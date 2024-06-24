package task

import "time"

type Task struct {
	Name    string     `json:"name"`
	Content string     `json:"content"`
	Until   *time.Time `json:"until,omitempty"`
}
