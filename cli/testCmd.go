package cli

import (
	"aoc-cli/aocweb"
	"aoc-cli/cli/flags"
	cli "aoc-cli/output"
	"aoc-cli/runner"
	"fmt"

	"github.com/spf13/cobra"
)

var testCommand = &cobra.Command{
	Use:   "test [task]",
	Short: "Test the given task (1 or 2) against the example data.",
	Long:  "Solves the given task (1 or 2) against the example data. \n Uses the specified day or current day. \n Uses the specified language or default language.",
	Run: func(cmd *cobra.Command, args []string) {
		day, year, lang, flagErr := flags.GetFlags(cmd)
		if flagErr != nil {
			cli.PrintFromError(flagErr).PrintError()
			return
		}

		task, taskErr := flags.GetTask(args)
		if taskErr != nil {
			cli.PrintFromError(taskErr).PrintError()
			return
		}

		cli.ToPrintf("Testing task %d of day %d in year %d using language %s", task, day, year, lang).PrintDebug()
		runResult := runner.TestDay(year, day, task, lang)
		expectedResult, err := aocweb.GetResource(fmt.Sprintf("testOutput%d", task), day, year)
		if err != nil {
			cli.ToPrint("Could not get solution for example data.").PrintWarning()
			cli.PrintFromError(err).PrintDebug()
			cli.ToPrintf("Your soulution: %s", runResult[len(runResult)-1]).PrintSuccess()
		} else if runResult[len(runResult)-1] == expectedResult {
			cli.ToPrint("Your solution is correct!").PrintSuccess()
			cli.ToPrintf("Expected: %s", expectedResult).PrintDebug()
			cli.ToPrintf("Got: %s", runResult[len(runResult)-1]).PrintDebug()
		} else {
			cli.ToPrintf("Your solution does not match the expected result.\n Expected: %s\n Got: %s", expectedResult, runResult[len(runResult)-1]).PrintError()
		}
	},
}

func init() {
	addCommand(testCommand)
	flags.AddPersistentFlags(testCommand)
}
