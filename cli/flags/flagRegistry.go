package flags

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var flagRegistry = pflag.NewFlagSet("FlagRegistry", pflag.PanicOnError)
var viperConfigToFlagName = map[string]string{}

func GetFlag(name string) *pflag.Flag {
	return flagRegistry.Lookup(name)
}

func GetConfigurableFlags() map[string]*pflag.Flag {
	result := map[string]*pflag.Flag{}

	for viperKey, flagName := range viperConfigToFlagName {
		result[viperKey] = flagRegistry.Lookup(flagName)
	}

	return result
}

func init() {
	currentTime := time.Now()

	flagRegistry.StringP("lang", "l", "", "Language to use")
	flagRegistry.IntP("day", "d", currentTime.Day(), "Day to use")
	flagRegistry.IntP("year", "y", currentTime.Year(), "Year to use")
	flagRegistry.Bool("debug", false, "Enable debug output")
	flagRegistry.Bool("no-emojis", false, "Disable emojis in the output")
	flagRegistry.StringP("cookie", "c", "", "Cookie for webrequests")
	flagRegistry.Bool("task2", false, "Include fetching the second challenge")
	flagRegistry.BoolP("submit", "s", false, "Submit the solution to the server")

	// Setting viper config names & entries
	viperConfigToFlagName["language"] = "lang"
	viperConfigToFlagName["noEmojis"] = "no-emojis"
	viperConfigToFlagName["cookie"] = "cookie"

	// Binding viper
	for viperKey, flagName := range viperConfigToFlagName {
		viper.BindPFlag(viperKey, flagRegistry.Lookup(flagName))
	}
}
