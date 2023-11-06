package cli

import (
	"aoc-cli/aocweb"
	cli "aoc-cli/output"
	"aoc-cli/runner"
	"fmt"

	"github.com/spf13/cobra"
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
		runResult := runner.SolveDay(year, day, task, lang)
		cli.PrintSuccessFmt("Your soulution: %s", runResult[len(runResult)-1])

		submit, err := cmd.Flags().GetBool("submit")
		if err != nil {
			cli.PrintDebug(err.Error())
			return
		}
		if submit {
			cli.PrintLog("Submitting solution")
			answer := aocweb.Submit(day, year, task, runResult[len(runResult)-1])
			if answer != nil {
				cli.PrintError(answer.Error())
			} else {
				cli.PrintSuccess("Your solution is correct!")
				cli.PrintSuccessFmt("Solved task %d of day %d of %d!", task, day, year)
			}
		}
	},
}

func init() {
	addCommand(solveCommand)
	addPersistentFlags(solveCommand)

	solveCommand.Flags().BoolP("submit", "s", false, "Submit the solution to the server")
}
