package cli

import (
	"aoc-cli/aocweb"
	cli "aoc-cli/output"
	"aoc-cli/utils"
	"errors"
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
			cli.PrintError(flagErr)
			return
		}

		cli.PrintDebugFmt("Initializing day %d in year %d using language %s", day, year, lang)
		p := cli.ProgressBar{}
		p.Run("Creating directories")
		dir := utils.GetChallengeDirectory(year, day)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			p.Cancel("Could not create directories")
			cli.PrintError(err)
			return
		}
		p.Set("Downloading resources", 0.1666)

		warningSent := false

		resourcesToGet := []string{"challenge1", "solveInput", "testInput", "testOutput1"}
		getSecond, err := cmd.Flags().GetBool("second")
		if err == nil && getSecond {
			resourcesToGet = append(resourcesToGet, "challenge2", "testOutput2")
		}

		for idx, resource := range resourcesToGet {
			_, err = aocweb.GetResource(resource, day, year)
			if err != nil {
				cli.PrintDebugFmt("Error requesting %s", resource)
				cli.PrintDebugError(err)
				if !warningSent {
					cli.PrintWarning("Could not access web page")
					warningSent = true
				}
			}
			p.Set("Downloading resources", (2.0+float64(idx))/(2.0+float64(len(resourcesToGet))))
		}

		p.Set("Initializing language", (1.0+float64(len(resourcesToGet)))/(2.0+float64(len(resourcesToGet))))

		filesToWrite := lang.GetFilesToWrite()
		for _, file := range filesToWrite {
			fullPath := fmt.Sprintf("%s%s", dir, file.Filename)
			if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
				cli.PrintDebugFmt("%s doesnt exist, writing template", fullPath)
				os.WriteFile(fullPath, []byte(file.Content), 0644)
			}
		}

		p.Finish(fmt.Sprintf("Day %d in year %d initialized", day, year))
	},
}

func init() {
	addCommand(initCommand)
	addPersistentFlags(initCommand)
	addCookieFlag(initCommand)

	initCommand.Flags().Bool("second", false, "Gets the second challange too.")
}
