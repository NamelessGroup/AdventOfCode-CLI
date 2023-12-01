package cli

import (
	"aoc-cli/cli/flags"
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
		validConfigs := flags.GetConfigurableFlags()

		if len(args) == 1 && args[0] == "list" {
			// List all valid options

			cli.ToPrint("Available config options:").PrintLog()
			for viperKey, flag := range validConfigs {
				cli.ToPrintf("    %s [%s] - %s", viperKey, flag.Value.Type(), flag.Usage).Print()
			}

			languageStr, err := cmd.Flags().GetString("lang")
			if err == nil && languageStr != "" {
				languageOptions := getLanguageSpecificConfigKeys(languageStr)
				if len(languageOptions) > 0 {
					cli.ToPrintf("Available language specific options for %s:", languageStr).PrintLog()
					for viperKey, flag := range languageOptions {
						cli.ToPrintf("    %s [%s] - %s", viperKey, flag.DataType, flag.Description).Print()
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

		languageSpecificOptions := map[string]utils.LanguageSpecificOption{}
		languageStr, err := cmd.Flags().GetString("lang")
		if err == nil {
			languageSpecificOptions = getLanguageSpecificConfigKeys(languageStr)
		}

		if slices.Contains(getConfigKeys(languageSpecificOptions), configName) {
			viper.Set(fmt.Sprintf("languages.%s.%s", languageStr, configName), configValue)
		} else if !slices.Contains(getConfigKeys(validConfigs), configName) {
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
	flags.AddConfigLanguageFlag(configCommand)
}

func getLanguageSpecificConfigKeys(lang string) map[string]utils.LanguageSpecificOption {
	if lang == "" {
		return map[string]utils.LanguageSpecificOption{}
	}

	languageObj, err := runner.ResolveLanguage(lang)
	if err != nil {
		cli.ToPrintf("Language %s not found", lang).PrintWarning()
		return map[string]utils.LanguageSpecificOption{}
	}
	return languageObj.GetLanguageSpecificConfigKeys()
}

func getConfigKeys[V interface{}](flags map[string]V) []string {
	result := []string{}

	for k := range flags {
		result = append(result, k)
	}

	return result
}
