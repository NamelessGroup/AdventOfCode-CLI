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
		validConfigs := getValidConfigs()

		if len(args) == 1 && args[0] == "list" {
			// List all valid options

			cli.ToPrint("Available config options:").PrintLog()
			for _, validCfg := range validConfigs {
				cli.ToPrintf("    %s [%s] - %s", validCfg.ViperKey, validCfg.DataType, validCfg.Description).Print()
			}

			languageStr, err := cmd.Flags().GetString("lang")
			if err == nil {
				languageOptions := getLanguageSpecificConfigKeys(languageStr)
				if len(languageOptions) > 0 {
					cli.ToPrintf("Available language specific options for %s:", languageStr).PrintLog()
					for _, validCfg := range languageOptions {
						cli.ToPrintf("    %s [%s] - %s", validCfg.ViperKey, validCfg.DataType, validCfg.Description).Print()
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

		languageSpecificOptions := map[string]utils.FlagMetadata{}
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
	addConfigLanguageFlag(configCommand)
}

func getValidConfigs() map[string]utils.FlagMetadata {
	result := map[string]utils.FlagMetadata{}

	for k := range Flags {
		if Flags[k].ViperKey != "" {
			result[k] = Flags[k]
		}
	}

	return result
}

func getLanguageSpecificConfigKeys(lang string) map[string]utils.FlagMetadata {
	if lang == "" {
		return map[string]utils.FlagMetadata{}
	}

	languageObj, err := runner.ResolveLanguage(lang)
	if err != nil {
		cli.ToPrintf("Language %s not found", lang).PrintWarning()
		return map[string]utils.FlagMetadata{}
	}
	return languageObj.GetLanguageSpecificConfigKeys()
}

func getConfigKeys(flags map[string]utils.FlagMetadata) []string {
	result := []string{}

	for k := range flags {
		result = append(result, flags[k].ViperKey)
	}

	return result
}
