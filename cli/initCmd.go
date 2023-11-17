package cli

import (
	"aoc-cli/aocweb"
	cli "aoc-cli/output"
	"aoc-cli/utils"
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
		dir := utils.GetChallengeDirectory(year, day)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			p.Cancel("Could not create directories")
			cli.PrintError(err.Error())
			return
		}
		p.Set("Downloading resources", 0.1666)

		warningSent := false

		for idx, resource := range []string{"challenge1", "solveInput", "testInput", "testOutput1"} {
			_, err = aocweb.GetResource(resource, day, year)
			if err != nil {
				cli.PrintDebugFmt("Error requesting %s", resource)
				cli.PrintDebug(err.Error())
				if !warningSent {
					cli.PrintWarning("Could not access web page")
					warningSent = true
				}
			}
			p.Set("Downloading resources", (2.0+float64(idx))/6.0)
		}

		p.Set("Initializing language", 0.8333)

		filesToWrite := lang.GetFilesToWrite()
		for _, file := range filesToWrite {
			os.WriteFile(fmt.Sprintf("%s/%s", dir, file.Filename), []byte(file.Content), 0644)
		}

		p.Finish(fmt.Sprintf("Day %d in year %d initialized", day, year))
	},
}

func init() {
	addCommand(initCommand)
	addPersistentFlags(initCommand)
}
