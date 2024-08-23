package main

import (
	"fmt"
	"hackerNewsApi/cmd/schedule/cron"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "batch",
	Short: "Command for start pub-sub redis",
	Long:  `Command for start pub-sub redis and schedule for cron api hn-item update items.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the root command is called
		fmt.Println("Welcome to batch processing on schedule job! Use --help for usage.")
	},
}

func init() {
	rootCmd.AddCommand(cron.ScheduleCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
