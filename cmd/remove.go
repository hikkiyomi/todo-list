package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hikkiyomi/todo/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes task of given id from todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		directoryPath := viper.Get("tasks.path").(string)
		id, err := cmd.Flags().GetInt("id")

		if err != nil {
			log.Fatal("Flag id is either not provided or not integer")
		}

		targetFile := fmt.Sprintf("%v%v.json", directoryPath, id)

		os.Remove(targetFile)
		util.Reorganize(directoryPath)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().Int("id", 0, "defines the id of removing task")

	removeCmd.MarkFlagRequired("id")
}
