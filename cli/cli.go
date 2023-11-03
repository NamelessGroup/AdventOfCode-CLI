package cli

import (
	cli "aoc-cli/output"
	"aoc-cli/runner"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "aoc-cli is a CLI for Advent of Code",
	Long:  "aoc-cli is a CLI for Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		cli.PrintWarning("No command specified")
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
		cli.PrintError(err.Error())
	}
}

func addCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

/*
Returns day, year, lang, error
*/
func getFlags(cmd *cobra.Command) (int, int, runner.Language, error) {
	day, dayErr := cmd.Flags().GetInt("day")
	year, yearErr := cmd.Flags().GetInt("year")
	if dayErr != nil {
		return -1, -1, nil, dayErr
	}
	if yearErr != nil {
		return -1, -1, nil, yearErr
	}

	currentTime := time.Now()
	if year == currentTime.Year() && day > currentTime.Day() {
		return -1, -1, nil, errors.New("Day must be before tomorrow")
	}
	if year < 2015 || year > currentTime.Year() {
		return -1, -1, nil, fmt.Errorf("Year must be between 2015 and %d", currentTime.Year())
	}
	if day < 1 || day > 25 {
		return -1, -1, nil, errors.New("Day must be between 1 and 25")
	}

	lang, langErr := cmd.Flags().GetString("lang")
	if langErr != nil {
		return -1, -1, nil, langErr
	}
	langResolved, err := runner.ResolveLanguage(lang)
	if err != nil {
		return -1, -1, nil, err
	}
	return day, year, langResolved, nil
}
