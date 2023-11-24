package cli

import (
	"aoc-cli/aocweb"
	cli "aoc-cli/output"
	"aoc-cli/runner"

	"github.com/spf13/cobra"
)

var solveCommand = &cobra.Command{
	Use:   "solve [task]",
	Short: "Solves the given task (1 or 2)",
	Long:  "Solves the given task (1 or 2). Defaults to 1. \n Uses the specified day or current day. \n Uses the specified language or default language.",
	Run: func(cmd *cobra.Command, args []string) {
		day, year, lang, flagErr := getFlags(cmd)
		if flagErr != nil {
			cli.PrintFromError(flagErr).PrintError()
			return
		}

		task, taskErr := getTask(args)
		if taskErr != nil {
			cli.PrintFromError(taskErr).PrintError()
			return
		}

		cli.ToPrintf("Solving task %d of day %d in year %d using language %s", task, day, year, lang).PrintDebug()
		runResult := runner.SolveDay(year, day, task, lang)
		cli.ToPrintf("Your soulution: %s", runResult[len(runResult)-1]).PrintSuccess()

		submit, err := cmd.Flags().GetBool("submit")
		if err != nil {
			cli.PrintFromError(err).PrintDebug()
			return
		}
		if submit {
			cli.ToPrint("Submitting solution").NewLine(false).PrintLog()
			answer := aocweb.Submit(day, year, task, runResult[len(runResult)-1])
			if answer != nil {
				cli.PrintFromError(answer).PrintError()
			} else {
				cli.ToPrint("Your solution is correct!").PrintSuccess()
				cli.ToPrintf("Solved task %d of day %d of %d!", task, day, year).PrintSuccess()
				_, err := aocweb.GetResource("challenge2", day, year)
				if err != nil {
					cli.ToPrint("Could not get 2nd challenge").PrintWarning()
					cli.PrintFromError(err).PrintDebug()
				}
			}
		}
	},
}

func init() {
	addCommand(solveCommand)
	addPersistentFlags(solveCommand)
	addCookieFlag(solveCommand)
	addSubmitFlag(solveCommand)
}
