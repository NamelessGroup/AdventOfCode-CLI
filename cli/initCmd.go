package cli

import (
	"aoc-cli/aocweb"
	"aoc-cli/output"
	"fmt"
	"os"
	"github.com/spf13/cobra"
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

		cli.PrintDebugFmt("Initializing day %d in year %d using language %s", day, year, lang)
		err := os.MkdirAll(fmt.Sprintf("%d/%d", year, day), 0755)
		if err != nil {
			cli.PrintError(err.Error())
			return
		}
		_, err = aocweb.GetResource("dayPage", day, year)
		if err != nil {
			cli.PrintWarning("Could not access web page")
			return
		}
	},
}

func init() {
	addCommand(initCommand)
	addPersistentFlags(initCommand)
}
