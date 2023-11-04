package runner

import (
	"aoc-cli/runner/languages"
	"fmt"
)

func ResolveLanguage(lang string) (Language, error) {
	languageMap := map[string]Language{
		"test": languages.Test{},
	}

	if languageMap[lang] == nil {
		return nil, fmt.Errorf("Language %s not found", lang)
	}

	return languageMap[lang], nil
}
