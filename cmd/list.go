/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hikkiyomi/todo/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows all your pending tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasksPath := viper.Get("tasks.path").(string)
		entries, err := os.ReadDir(tasksPath)

		if err != nil {
			log.Fatal(err)
		}

		tasks := make([]task.Task, 0)

		for _, entry := range entries {
			if strings.HasSuffix(entry.Name(), ".json") {
				path := fmt.Sprintf("%v%v", tasksPath, entry.Name())
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

		result := "=========================\n"

		for _, task := range tasks {
			result += task.String() + "\n=========================\n"
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
