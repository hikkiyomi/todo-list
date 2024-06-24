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
	Short: "Adds your tasks to todo-list",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		content := args[1]
		untilString, _ := cmd.Flags().GetString("until")
		untilTime, err := time.Parse("2006-01-02 15:04:05", untilString)

		if err != nil {
			fmt.Println("Could not convert your data into normal format.")
			return
		}

		task := task.Task{Name: name, Content: content, Until: untilTime}
		taskJson, err := json.Marshal(task)

		if err != nil {
			fmt.Println("Could not serialize your task into json.")
			return
		}

		os.WriteFile("task.json", taskJson, 0644)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("until", "", "defines the deadline of the task")
}
