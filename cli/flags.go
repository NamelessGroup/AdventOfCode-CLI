package cli

import (
	"aoc-cli/runner"
	"aoc-cli/utils"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addPersistentFlags(cmd *cobra.Command) {
	currentTime := time.Now()

	cmd.PersistentFlags().StringP("lang", "l", viper.GetString("lang"), "Language to run")
	cmd.PersistentFlags().IntP("day", "d", currentTime.Day(), "Day to run")
	cmd.PersistentFlags().IntP("year", "y", currentTime.Year(), "Year to run")
	cmd.PersistentFlags().Bool("debug", false, "Enable debug output")
}

func addCookieFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("cookie", "c", viper.GetString("cookie"), "Cookie for web requests")
	viper.BindPFlag("cookie", solveCommand.Flags().Lookup("cookie"))
}

func addSecondChallengeFlag(cmd *cobra.Command) {
	initCommand.Flags().Bool("second", false, "Include fetching the second challange")
}

func addSubmitFlag(cmd *cobra.Command) {
	solveCommand.Flags().BoolP("submit", "s", false, "Submit the solution to the server")
}

func getTask(args []string) (int, error) {
	if len(args) == 0 {
		return 1, nil
	}

	task, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, err
	}

	if task != 1 && task != 2 {
		return 0, utils.AOCCLIError("Task must be 1 or 2").DebugInfof("runnerCmds", "Supplied task: %d", task)
	}
	return task, nil
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
