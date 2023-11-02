package cli

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "aoc-cli is a CLI for Advent of Code",
	Long:  "aoc-cli is a CLI for Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		PrintWarning("No command specified")
	},
}

func addPersistentFlags(cmd *cobra.Command) {
	currentTime := time.Now()

	cmd.PersistentFlags().StringP("lang", "l", "python", "Language to run")
	cmd.PersistentFlags().IntP("day", "d", currentTime.Day(), "Day to run")
	cmd.PersistentFlags().IntP("year", "y", currentTime.Year(), "Year to run")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		PrintError(err.Error())
	}
}

func addCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

/*
Returns day, year, lang, error
*/
func getFlags(cmd *cobra.Command) (int, int, string, error) {
	day, dayErr := cmd.Flags().GetInt("day")
	year, yearErr := cmd.Flags().GetInt("year") 
	if dayErr != nil {
		return -1, -1, "", dayErr
	}
	if yearErr != nil {
		return -1, -1, "", yearErr
	}

	currentTime := time.Now()
	if year == currentTime.Year() && day > currentTime.Day() {
		return -1, -1, "", errors.New("Day must be before tomorrow")
	}
	if year < 2015 || year > currentTime.Year() {
		return -1, -1, "", fmt.Errorf("Year must be between 2015 and %d", currentTime.Year())
	}
	if day < 1 || day > 25  {
		return -1, -1, "", errors.New("Day must be between 1 and 25")
	}

	lang, langErr := cmd.Flags().GetString("lang")
	if langErr != nil {
		return -1, -1, "", langErr
	}
	if lang != "python" && lang != "go" {
		return -1, -1, "", errors.New("Unknown language")
	}
	return day, year, lang, nil
}
