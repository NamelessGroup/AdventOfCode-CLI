package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"aoc-cli/output"
)

var initCommand = &cobra.Command{
	Use:   "init [day]",
	Short: "Sets up the files for a given day",
	Long:  "Sets up the files for a given day. \n If no day is specified it uses the current day. \n If no language is specified it uses your default language.",
	Run: func(cmd *cobra.Command, args []string) {
		day, year, lang, flagErr := getFlags(cmd)
		if flagErr != nil {
			cli.PrintError(flagErr.Error())
			return
		}

		cli.PrintDebug(fmt.Sprintf("Initializing day %d in year %d using language %s", day, year, lang))
		
	},
}

func init() {
	addCommand(initCommand)
	addPersistentFlags(initCommand)
}
