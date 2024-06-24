package task

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

func (task *Task) Parse(flags *pflag.FlagSet) {
	name, _ := flags.GetString("name")
	content, _ := flags.GetString("content")
	task.Name, task.Content = name, content

	untilString, _ := flags.GetString("until")
	untilTime, err := time.Parse("2006-01-02 15:04", untilString)

	if err == nil {
		task.Until = &untilTime
	}
}

func (task *Task) GetTaskPath() string {
	directoryPath := viper.Get("tasks.path")

	return fmt.Sprintf("%v%v.json", directoryPath, task.Name)
}
