package cli

import (
	"aoc-cli/aocweb"
	cli "aoc-cli/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "aoc-cli is a CLI for Advent of Code",
	Long:  "aoc-cli is a CLI for Advent of Code",
	Run: func(cmd *cobra.Command, args []string) {
		cli.PrintWarning("No command specified")
		_, err := aocweb.GetResource("challenge2", 3, 2022)
		println(err.Error())
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		cli.PrintDebugMessages = debug
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cli.PrintError(err)
	}
}

func init() {
	viper.SetConfigName("aoc-cli-config")
	viper.AddConfigPath("$HOME/.config/aoc-cli")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	viper.SetDefault("lang", "test")

	err := viper.ReadInConfig()
	if err != nil {
		cli.PrintWarning("Could not read config file")
	}
}

func addCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
