package cli

import (
	"aoc-cli/output"
	"aoc-cli/runner"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var solveCommand = &cobra.Command{
	Use:   "solve [task]",
	Short: "Solves the given task (1 or 2)",
	Long:  "Solves the given task (1 or 2). Defaults to 1. \n Uses the specified day or current day. \n Uses the specified language or default language.",
	Run: func(cmd *cobra.Command, args []string) {
		day, year, lang, flagErr := getFlags(cmd)
		if flagErr != nil {
			cli.PrintError(flagErr.Error())
			return
		}

		task, taskErr := getTask(args)
		if taskErr != nil {
			cli.PrintError(taskErr.Error())
			return
		}

		cli.PrintDebug(fmt.Sprintf("Solving task %d of day %d in year %d using language %s", task, day, year, lang))
		runner.SolveDay(day, task, lang)
	},
}

func init() {
	addCommand(solveCommand)
	addPersistentFlags(solveCommand)

	solveCommand.Flags().StringP("cookie", "c", viper.GetString("cookie"), "Session cookie to use for web requests")
	viper.BindPFlag("cookie", solveCommand.Flags().Lookup("cookie"))
}
