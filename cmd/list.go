package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/hikkiyomi/todo/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getAllTasks() []task.Task {
	directoryPath := viper.Get("tasks.path").(string)
	entries, err := os.ReadDir(directoryPath)

	if err != nil {
		log.Fatal(err)
	}

	tasks := make([]task.Task, 0, cap(entries))

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".json") {
			path := fmt.Sprintf("%v%v", directoryPath, entry.Name())
			bytes, err := os.ReadFile(path)

			if err != nil {
				log.Fatal("Could not read file located at " + path)
			}

			var task task.Task

			err = json.Unmarshal(bytes, &task)

			if err != nil {
				log.Fatal("Could not deserialize json into task.")
			}

			tasks = append(tasks, task)
		}
	}

	return tasks
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows all your pending tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := getAllTasks()
		table := table.New(os.Stdout)

		table.SetHeaders("ID", "Name", "Content", "Until")

		for i, task := range tasks {
			var untilString string

			if task.Until == nil {
				untilString = ""
			} else {
				untilString = task.Until.String()
			}

			table.AddRow(strconv.Itoa(i+1), task.Name, task.Content, untilString)
		}

		table.Render()
		fmt.Println("Use `todo info <taskId>` to get more info about task.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
