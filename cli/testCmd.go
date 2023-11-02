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
		fmt.Println("solve")
	},
}

func init() {
	AddCommand(testCommand)
	AddPersistentFlags(testCommand)
}
