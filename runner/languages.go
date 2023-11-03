package runner

import (
	cli "aoc-cli/output"
	"aoc-cli/runner/languages"
)

func ResolveLanguage(lang string) Language {
	languageMap := map[string]Language{
		"test": languages.Test{},
	}

	if languageMap[lang] == nil {
		cli.PrintErrorFmt("Error resolving language %s", lang)
	}

	return languageMap[lang]
}
