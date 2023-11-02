package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testCommand = &cobra.Command{
	Use:   "solve [task]",
	Short: "Test the given task (1 or 2) against the example data.",
	Long:  "Solves the given task (1 or 2) against the example data. \n Uses the specified day or current day. \n Uses the specified language or default language.",
	Run: func(cmd *cobra.Command, args []string) {
		day, year, lang, flagErr := getFlags(cmd)
		if flagErr != nil {
			PrintError(flagErr.Error())
			return
		}

		task, taskErr := getTask(args)
		if taskErr != nil {
			PrintError(taskErr.Error())
			return
		}

		print(fmt.Sprintf("Testing task %d of day %d in year %d using language %s", task, day, year, lang))
	},
}

func init() {
	addCommand(testCommand)
	addPersistentFlags(testCommand)
}
