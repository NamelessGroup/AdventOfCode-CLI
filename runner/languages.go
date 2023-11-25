package runner

import (
	"aoc-cli/runner/languages"
	"aoc-cli/utils"
)

type Language interface {
	GetSolveCommand(directoryPath string, task int) utils.ExecutionDetails
	GetTestCommand(directoryPath string, task int) utils.ExecutionDetails
	GetPreparationCommand(directoryPath string, task int) []utils.ExecutionDetails
	GetFilesToWrite() []utils.FileTemplate
	GetLanguageSpecificConfigKeys() map[string]utils.FlagMetadata
}

func ResolveLanguage(lang string) (Language, error) {
	languageMap := map[string]Language{
		"python":     languages.Python{},
		"java":       languages.Java{},
		"typescript": languages.TypeScript{},
		"haskell":    languages.Haskell{},
	}

	if languageMap[lang] == nil {
		return nil, utils.AOCCLIErrorf("Language %s not found", lang)
	}

	return languageMap[lang], nil
}
