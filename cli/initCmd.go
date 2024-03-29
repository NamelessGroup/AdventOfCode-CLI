package cli

import (
	"aoc-cli/aocweb"
	"aoc-cli/cli/flags"
	cli "aoc-cli/output"
	"aoc-cli/utils"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Sets up the files for a given day",
	Long:  "Sets up the files for a given day. \n If no day is specified it uses the current day. \n If no language is specified it uses your default language.",
	Run: func(cmd *cobra.Command, args []string) {
		day, year, lang, flagErr := flags.GetFlags(cmd)
		if flagErr != nil {
			cli.PrintFromError(flagErr).PrintError()
			return
		}

		cli.ToPrintf("Initializing day %d in year %d using language %s", day, year, lang).PrintDebug()

		resourcesToGet := []string{"challenge1", "solveInput", "testInput", "testOutput1"}
		getSecond, err := cmd.Flags().GetBool("task2")
		if err == nil && getSecond {
			resourcesToGet = append(resourcesToGet, "challenge2", "testOutput2")
		}

		p := cli.NewProgressBar(2+len(resourcesToGet), "Creating directories")

		dir := utils.GetChallengeDirectory(year, day)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			p.Cancel("Could not create directories")
			cli.PrintFromError(err).PrintError()
			return
		}
		p.GotoNextTask("Downloading resources")

		warningSent := false

		for _, resource := range resourcesToGet {
			_, err = aocweb.GetResource(resource, day, year)
			if err != nil {
				cli.ToPrintf("Error requesting %s", resource).PrintDebug()
				cli.PrintFromError(err).PrintDebug()
				if !warningSent {
					cli.ToPrint("Could not access web page").PrintWarning()
					warningSent = true
				}
			}
			p.GotoNextTask("Downloading ressources")
		}

		p.GotoNextTask("Initializing language")

		filesToWrite := lang.GetFilesToWrite()
		for _, file := range filesToWrite {
			fullPath := fmt.Sprintf("%s%s", dir, file.Filename)
			if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
				cli.ToPrintf("%s doesnt exist, writing template", fullPath).PrintDebug()
				os.WriteFile(fullPath, []byte(file.Content), 0644)
			}
		}

		p.Finish(fmt.Sprintf("Day %d in year %d initialized", day, year))
	},
}

func init() {
	addCommand(initCommand)
	flags.AddPersistentFlags(initCommand)
	flags.AddCookieFlag(initCommand)
	flags.AddSecondChallengeFlag(initCommand)
}
