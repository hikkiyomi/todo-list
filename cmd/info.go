package cmd

import (
	"fmt"
	"log"

	"github.com/hikkiyomi/todo/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gives more info about task.",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")

		if err != nil {
			log.Fatal(err)
		}

		var task task.Task

		directoryPath := viper.Get("tasks.path")
		path := fmt.Sprintf("%v%v.json", directoryPath, id)

		task.Read(path)
		fmt.Println(task.String())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().Int("id", 0, "defines the id of requested task")
	infoCmd.MarkFlagRequired("id")
}
