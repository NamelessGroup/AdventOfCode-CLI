package aocweb

import (
	cli "aoc-cli/output"
	"bytes"
	"fmt"

	"net/http"
	"regexp"

	"github.com/spf13/viper"
)

func get(day int, year int, path string) (string, error) {
	cli.PrintDebugFmt("Getting %s for day %d year %d from the website", path, day, year)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d%s", year, day, path), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", viper.GetString("cookie")))
	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	cli.PrintDebugFmt("Got status code %d", result.StatusCode)
	if result.StatusCode != 200 {
		return "", fmt.Errorf("Got status code %d", result.StatusCode)
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(result.Body)
	body := buffer.String() 

	return string(body), nil
}

func GetDayPage(day int, year int) (string, error) {
	html, err := get(day, year, "")
	if err != nil {
		return "", err
	}
	wholeArticle := regexp.MustCompile("(?ms)<article(.*)</article>").FindString(html)
	wholeArticle = replaceTagRegex(wholeArticle, "</?article>", "")

	wholeArticle = replaceTagRegex(wholeArticle, "<pre><code>", "\n```\n")
	wholeArticle = replaceTagRegex(wholeArticle, "</code></pre>", "```")

	wholeArticle = replaceTagRegex(wholeArticle, "<code><em>", "**`")
	wholeArticle = replaceTagRegex(wholeArticle, "</em></code>", "`**")

	wholeArticle = replaceTagRegex(wholeArticle, "<code>", "`")
	wholeArticle = replaceTagRegex(wholeArticle, "</code>", "`")

	wholeArticle = replaceTagRegex(wholeArticle, "<em>", "**")
	wholeArticle = replaceTagRegex(wholeArticle, "</em>", "**")

	wholeArticle = replaceTagRegex(wholeArticle, "<p>", "")
	wholeArticle = replaceTagRegex(wholeArticle, "</p>", "\n")

	wholeArticle = replaceTagRegex(wholeArticle, "<h2>---", "<h2>")
	wholeArticle = replaceTagRegex(wholeArticle, "<h2>", "#")
	wholeArticle = replaceTagRegex(wholeArticle, "---</h2>", "</h2>")
	wholeArticle = replaceTagRegex(wholeArticle, "</h2>", "\n")

	wholeArticle = replaceTagRegex(wholeArticle, "</?ul>", "")
	wholeArticle = replaceTagRegex(wholeArticle, "<li>", "- ")
	wholeArticle = replaceTagRegex(wholeArticle, "</li>", "")

	// replace links
	wholeArticle = regexp.MustCompile("<a href=\"(.*)\">(.*)</a>").ReplaceAllStringFunc(wholeArticle, func(match string) string {
		link := regexp.MustCompile("<a href=\"([^\"]*)\"[^>]*>").FindStringSubmatch(match)[1]
		text := regexp.MustCompile("<a[^>]*>(.*)</a>").FindStringSubmatch(match)[1]
		return fmt.Sprintf("[%s](%s)", text, link)
	})

	// eliminate all other tags
	wholeArticle = replaceTagRegex(wholeArticle, "<>", "")
	return wholeArticle, nil
}

func replaceTagRegex(text string, tagRegex string, replacement string) string {
	replacesRegex := regexp.MustCompile(">").ReplaceAllString(tagRegex, "[^>]*>")
	return regexp.MustCompile(replacesRegex).ReplaceAllString(text, replacement)
}

func GetSolveInput(day int, year int) (string, error) {
	return get(day, year, "/input")
}

func GetTestInput(day int, year int) (string, error) {
	dayPage, err := GetResource("challenge", day, year)
	if err != nil {
		return "", err
	}
	return regexp.MustCompile("```([^`]*)```").FindStringSubmatch(dayPage)[1], nil	
}

