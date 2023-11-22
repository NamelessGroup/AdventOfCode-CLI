package cli

import (
	cli "aoc-cli/output"
	"aoc-cli/utils"
	"errors"
	"os"
	"slices"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCommand = &cobra.Command{
	Use:   "config [key] [value]",
	Short: "Sets the config in the config file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		if strings.Trim(viper.GetViper().ConfigFileUsed(), " ") == "" {
			configName := "aoc-cli-config.json"
			if _, err := os.Stat(configName); errors.Is(err, os.ErrNotExist) {
				cli.PrintLog("Writing config file", true)
				os.WriteFile(configName, []byte("{}"), 0644)
			}
			err := viper.ReadInConfig()
			if err != nil {
				cli.PrintError(utils.AOCCLIError("Could not create config file"))
			}
		}

		validConfigs := []string{"cookie", "lang"}
		configName := args[0]
		configValue := args[1]


		if !slices.Contains(validConfigs, configName) {
			cli.PrintError(utils.AOCCLIErrorf("%s is not a valid config point", configName))
		}
		
		viper.Set(configName, configValue)
		err := viper.WriteConfig()
		if err != nil {
			cli.PrintError(err)
		}
	},
}

func init() {
	addCommand(configCommand)
}
