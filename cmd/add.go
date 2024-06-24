package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/hikkiyomi/todo-list/internal/task"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds your tasks to todo list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		content, _ := cmd.Flags().GetString("content")
		task := task.Task{Name: name, Content: content}

		untilString, _ := cmd.Flags().GetString("until")
		untilTime, err := time.Parse("2006-01-02 15:04", untilString)

		if err == nil {
			task.Until = &untilTime
		}

		taskJson, err := json.Marshal(task)

		if err != nil {
			fmt.Println("Could not serialize your task into json.")

			return
		}

		os.WriteFile(fmt.Sprintf("~/todos/%v.json", name), taskJson, 0644)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("name", "", "defines the name of the task")
	addCmd.Flags().String("content", "", "defines the content of the task")
	addCmd.Flags().String("until", "", "defines the deadline of the task")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("content")
}
