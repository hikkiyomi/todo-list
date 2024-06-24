package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hikkiyomi/todo/internal/task"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds your tasks to todo list",
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
			log.Fatal("Could not serialize your task into json.")
		}

		tasksPath := viper.Get("tasks.path")
		resultPath := fmt.Sprintf("%v%v.json", tasksPath, name)

		os.WriteFile(resultPath, taskJson, 0644)
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
