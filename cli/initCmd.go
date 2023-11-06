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
		p := cli.ProgressBar{}
		p.Run("Creating directories")
		err := os.MkdirAll(fmt.Sprintf("%d/%d", year, day), 0755)
		if err != nil {
			p.Cancel("Could not create directories")
			cli.PrintError(err.Error())
			return
		}
		p.Set("Downloading resources", 0.3333)
		_, err = aocweb.GetResource("challenge", day, year)
		if err != nil {
			cli.PrintDebug(err.Error())
			cli.PrintWarning("Could not access web page")
		}
		p.Set("Initializing language", 0.6666)
		// TODO: Create language specific files
		p.Finish(fmt.Sprintf("Day %d in year %d initialized", day, year))
	},
}

func init() {
	addCommand(initCommand)
	addPersistentFlags(initCommand)
}
