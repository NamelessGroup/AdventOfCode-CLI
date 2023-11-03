package runner

import (
	"aoc-cli/runner/languages"
	"errors"
)

func ResolveLanguage(lang string) (Language, error) {
	languageMap := map[string]Language{
		"test": languages.Test{},
	}

	if languageMap[lang] == nil {
		return nil, errors.New("Language not found")
	}

	return languageMap[lang], nil
}
