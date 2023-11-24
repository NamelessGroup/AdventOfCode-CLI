package cli

import (
	cli "aoc-cli/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "aoc-cli is a CLI for Advent of Code",
	Long:  "aoc-cli is a CLI for Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		cli.ToPrint("No command specified").PrintWarning()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		disableEmojis := viper.GetBool("disableEmojis")
		cli.PrintDebugMessages = debug
		cli.DisableEmojis = disableEmojis
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cli.PrintFromError(err).PrintError()
	}
}

func init() {
	viper.SetConfigName("aoc-cli-config")
	viper.AddConfigPath("$HOME/.config/aoc-cli")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		cli.ToPrint("Could not read config file").PrintWarning()
	}
}

func addCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
