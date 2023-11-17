package cli

import (
	cli "aoc-cli/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCommand = &cobra.Command{
	Use:   "config [key] [value]",
	Short: "Sets up the files for a given day",
	Long:  "Sets up the files for a given day. \n If no day is specified it uses the current day. \n If no language is specified it uses your default language.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(args[0], args[1])
		err := viper.WriteConfig()
		if err != nil {
			cli.PrintError(err)
			return
		}
	},
}

func init() {
	addCommand(configCommand)
}
