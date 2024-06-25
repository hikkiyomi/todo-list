package cmd

import (
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
			var task task.Task
			path := fmt.Sprintf("%v%v", directoryPath, entry.Name())

			task.Read(path)
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

		if len(tasks) == 0 {
			fmt.Println("No pending tasks.")

			return
		}

		table := table.New(os.Stdout)
		table.SetHeaders("ID", "Short", "Until")

		for i, task := range tasks {
			var untilString string

			if task.Until == nil {
				untilString = ""
			} else {
				untilString = task.Until.String()
			}

			table.AddRow(strconv.Itoa(i+1), task.Short, untilString)
		}

		table.Render()
		fmt.Println("Use `todo info <taskId>` to get more info about task.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
