package cli

import (
	"aoc-cli/aocweb"
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
		day, year, lang, flagErr := getFlags(cmd)
		if flagErr != nil {
			cli.PrintError(flagErr)
			return
		}

		task, taskErr := getTask(args)
		if taskErr != nil {
			cli.PrintError(taskErr)
			return
		}

		cli.PrintDebug(fmt.Sprintf("Testing task %d of day %d in year %d using language %s", task, day, year, lang))
		runResult := runner.TestDay(year, day, task, lang)
		expectedResult, err := aocweb.GetResource(fmt.Sprintf("testOutput%d", task), year, day)
		if err != nil {
			cli.PrintWarning("Could not get solution for example data.")
			cli.PrintDebugError(err)
			cli.PrintSuccessFmt("Your soulution: %s", runResult[len(runResult)-1])
		} else if runResult[len(runResult)-1] == expectedResult {
			cli.PrintSuccess("Your solution is correct!")
			cli.PrintDebug(fmt.Sprintf("Expected: %s", expectedResult))
			cli.PrintDebug(fmt.Sprintf("Got: %s", runResult[len(runResult)-1]))
		} else {
			cli.PrintErrorFmt("Your solution does not match the expected result.\n Expected: %s\n Got: %s", expectedResult, runResult[len(runResult)-1])
		}
	},
}

func init() {
	addCommand(testCommand)
	addPersistentFlags(testCommand)
}
