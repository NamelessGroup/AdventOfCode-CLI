package cli

import (
	cli "aoc-cli/output"
	"aoc-cli/runner"
	"aoc-cli/utils"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "aoc-cli is a CLI for Advent of Code",
	Long:  "aoc-cli is a CLI for Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		cli.PrintWarning("No command specified")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		cli.PrintDebugMessages = debug
	},
}

func addPersistentFlags(cmd *cobra.Command) {
	currentTime := time.Now()

	cmd.PersistentFlags().StringP("lang", "l", viper.GetString("lang"), "Language to run")
	cmd.PersistentFlags().IntP("day", "d", currentTime.Day(), "Day to run")
	cmd.PersistentFlags().IntP("year", "y", currentTime.Year(), "Year to run")
	cmd.PersistentFlags().Bool("debug", false, "Enable debug output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cli.PrintError(err)
	}
}

func init() {
	viper.SetConfigName("aoc-cli-config")
	viper.AddConfigPath("$HOME/.config/aoc-cli")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	viper.SetDefault("lang", "test")

	err := viper.ReadInConfig()
	if err != nil {
		cli.PrintWarning("Could not read config file")
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
		return -1, -1, nil, utils.AOCCLIError("Day must be before tomorrow").DebugInfof("cli", "Inputted day: %d", day)
	}
	if year < 2015 || year > currentTime.Year() {
		return -1, -1, nil, utils.AOCCLIErrorf("Year must be between 2015 and %d", currentTime.Year()).DebugInfof("cli", "Inputted year: %d", year)
	}
	if day < 1 || day > 25 {
		return -1, -1, nil, utils.AOCCLIError("Day must be between 1 and 25").DebugInfof("cli", "Inputted day: %d", day)
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
