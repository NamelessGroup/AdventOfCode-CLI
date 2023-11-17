package runner

import (
	"aoc-cli/runner/languages"
	"aoc-cli/utils"
)

func ResolveLanguage(lang string) (Language, error) {
	languageMap := map[string]Language{
		"test":   languages.Test{},
		"python": languages.Python{},
	}

	if languageMap[lang] == nil {
		return nil, utils.AOCCLIErrorf("Language %s not found", lang)
	}

	return languageMap[lang], nil
}
