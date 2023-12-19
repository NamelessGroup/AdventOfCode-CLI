package flags

import (
	"aoc-cli/runner"
	"aoc-cli/utils"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().AddFlag(GetFlag("lang"))
	cmd.PersistentFlags().AddFlag(GetFlag("day"))
	cmd.PersistentFlags().AddFlag(GetFlag("year"))
	cmd.PersistentFlags().AddFlag(GetFlag("debug"))
	cmd.PersistentFlags().AddFlag(GetFlag("no-emojis"))
}

func AddCookieFlag(cmd *cobra.Command) {
	cmd.Flags().AddFlag(GetFlag("cookie"))
}

func AddSecondChallengeFlag(cmd *cobra.Command) {
	cmd.Flags().AddFlag(GetFlag("task2"))
}

func AddSubmitFlag(cmd *cobra.Command) {
	cmd.Flags().AddFlag(GetFlag("submit"))
}

func AddConfigLanguageFlag(cmd *cobra.Command) {
	cmd.Flags().AddFlag(GetFlag("lang"))
}

func GetTask(args []string) (int, error) {
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
func GetFlags(cmd *cobra.Command) (int, int, runner.Language, error) {
	day, dayErr := cmd.Flags().GetInt("day")
	year := viper.GetInt("year")
	if dayErr != nil {
		return -1, -1, nil, dayErr
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

	lang := viper.GetString("language")
	langResolved, err := runner.ResolveLanguage(lang)
	if err != nil {
		return -1, -1, nil, err
	}
	return day, year, langResolved, nil
}
