package task

import (
	"fmt"
	"time"
)

type Task struct {
	Name    string     `json:"name"`
	Content string     `json:"content"`
	Until   *time.Time `json:"until,omitempty"`
}

func (task *Task) String() (result string) {
	result += fmt.Sprintf("NAME: %s\n\n", task.Name)
	result += fmt.Sprintf("CONTENT:\n%s\n\n", task.Content)

	if task.Until != nil {
		result += fmt.Sprintf("UNTIL: %s", task.Until)
	} else {
		result += "UNTIL: No deadline"
	}

	return
}
