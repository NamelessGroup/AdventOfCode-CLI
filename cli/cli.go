package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "aoc-cli is a CLI for Advent of Code",
	Long:  "aoc-cli is a CLI for Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root")
	},
}

func AddPersistentFlags(cmd *cobra.Command) {
	currentTime := time.Now()

	cmd.PersistentFlags().StringP("lang", "l", "python", "Language to run")
	cmd.PersistentFlags().IntP("day", "d", currentTime.Day(), "Day to run")
	cmd.PersistentFlags().IntP("year", "y", currentTime.Year(), "Year to run")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
