package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var solveCommand = &cobra.Command{
	Use:   "solve [task]",
	Short: "Solves the given task (1 or 2)",
	Long:  "Solves the given task (1 or 2). \n Uses the specified day or current day. \n Uses the specified language or default language.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solve")
	},
}

func init() {
	AddCommand(solveCommand)
	AddPersistentFlags(solveCommand)
}
