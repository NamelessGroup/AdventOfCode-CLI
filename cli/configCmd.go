package cli

import (
	cli "aoc-cli/output"
	"aoc-cli/runner"
	"aoc-cli/utils"
	"errors"
	"fmt"
	"os"
	"slices"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCommand = &cobra.Command{
	Use:   "config [key]|list [value]",
	Short: "Sets the config in the config file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		validConfigs := []string{"cookie", "language", "disableEmojis"}

		if len(args) == 1 && args[0] == "list" {
			// List all valid options

			cli.ToPrint("Available config options:").PrintLog()
			for _, validCfg := range validConfigs {
				cli.ToPrintf("    %s", validCfg).Print()
			}

			languageStr, err := cmd.Flags().GetString("lang")
			if err == nil && languageStr != "" {
				languageObj, err := runner.ResolveLanguage(languageStr)
				if err != nil {
					cli.ToPrintf("Language %s not found", languageStr).PrintWarning()
					return
				}
				languageOptions := languageObj.GetLanguageSpecificConfigKeys()
				if len(languageOptions) > 0 {
					cli.ToPrintf("Available language specific options for %s:", languageStr).PrintLog()
					for _, validCfg := range languageOptions {
						cli.ToPrintf("    %s", validCfg).Print()
					}
				} else {
					cli.ToPrintf("Language %s has no specific options.", languageStr).PrintLog()
				}
			}

			return
		} else if len(args) != 2 {
			cli.ToPrint("Please specify a key and value!").PrintError()
			return
		}

		if strings.Trim(viper.GetViper().ConfigFileUsed(), " ") == "" {
			configName := "aoc-cli-config.json"
			if _, err := os.Stat(configName); errors.Is(err, os.ErrNotExist) {
				cli.ToPrint("Writing config file").PrintLog()
				os.WriteFile(configName, []byte("{}"), 0644)
			}
			err := viper.ReadInConfig()
			if err != nil {
				cli.PrintFromError(utils.AOCCLIError("Could not create config file")).PrintError()
			}
		}

		configName := args[0]
		configValue := args[1]

		languageSpecificOptions := []string{}
		languageStr, err := cmd.Flags().GetString("lang")
		if err == nil && languageStr != "" {
			languageObj, err := runner.ResolveLanguage(languageStr)
			if err != nil {
				cli.ToPrintf("Language %s not found", languageStr).PrintWarning()
				return
			}
			languageSpecificOptions = languageObj.GetLanguageSpecificConfigKeys()
		}

		if slices.Contains(languageSpecificOptions, configName) {
			viper.Set(fmt.Sprintf("languages.%s.%s", languageStr, configName), configValue)
		} else if !slices.Contains(validConfigs, configName) {
			cli.PrintFromError(utils.AOCCLIErrorf("%s is not a valid config point", configName)).PrintError()
			return
		}

		viper.Set(configName, configValue)
		err = viper.WriteConfig()
		if err != nil {
			cli.PrintFromError(err).PrintError()
		}
	},
}

func init() {
	addCommand(configCommand)
	addConfigLanguageFlag(configCommand)
}
