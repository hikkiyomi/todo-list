package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-list",
	Short: "A small application for tracking your tasks.",
	Long:  ``,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
