package cli

import (
	"aoc-cli/runner"
	"aoc-cli/utils"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// This is to make viper work correctly, if the same flag is defined multiple times (e.g. in different commands)
var (
	_langFlag          string
	_disableEmojisFlag bool
	_cookieFlag        string
)

func addPersistentFlags(cmd *cobra.Command) {
	currentTime := time.Now()

	cmd.PersistentFlags().StringVarP(&_langFlag, "lang", "l", "", "Language to run")
	viper.BindPFlag("language", cmd.PersistentFlags().Lookup("lang"))
	cmd.PersistentFlags().IntP("day", "d", currentTime.Day(), "Day to run")
	cmd.PersistentFlags().IntP("year", "y", currentTime.Year(), "Year to run")
	cmd.PersistentFlags().Bool("debug", false, "Enable debug output")

	cmd.PersistentFlags().BoolVar(&_disableEmojisFlag, "disable-emojis", false, "Disable emojis in the output")
	viper.BindPFlag("disableEmojis", cmd.PersistentFlags().Lookup("disable-emojis"))
}

func addCookieFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&_cookieFlag, "cookie", "c", "", "Cookie for web requests")
	viper.BindPFlag("cookie", cmd.Flags().Lookup("cookie"))
}

func addSecondChallengeFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("second", false, "Include fetching the second challange")
}

func addSubmitFlag(cmd *cobra.Command) {
	cmd.Flags().BoolP("submit", "s", false, "Submit the solution to the server")
}

func addConfigLanguageFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("lang", "l", "", "Language to list options from")
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

	lang := viper.GetString("language")
	langResolved, err := runner.ResolveLanguage(lang)
	if err != nil {
		return -1, -1, nil, err
	}
	return day, year, langResolved, nil
}
