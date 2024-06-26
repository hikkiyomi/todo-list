package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/hikkiyomi/todo/internal/task"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds your tasks to todo list",
	Run: func(cmd *cobra.Command, args []string) {
		var task task.Task

		task.Parse(cmd.Flags())
		taskJson, err := json.Marshal(task)

		if err != nil {
			log.Fatal("Could not serialize your task into json.")
		}

		os.WriteFile(task.GetTaskPath(), taskJson, 0644)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("short", "", "defines the short description of the task")
	addCmd.Flags().String("long", "", "defines the long description of the task")
	addCmd.Flags().String("until", "", "defines the deadline of the task")

	addCmd.MarkFlagRequired("short")
}
