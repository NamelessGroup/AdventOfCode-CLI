package aocweb

import (
	cli "aoc-cli/output"
	"aoc-cli/utils"
	"bytes"
	"fmt"
	"strings"
	"net/http"
	"regexp"

	"github.com/spf13/viper"
	"html"
)

func get(day int, year int, path string) (string, error) {
	cli.ToPrintf("Getting %s for day %d year %d from the website", path, day, year).PrintDebug()
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d%s", year, day, path), nil)
	if err != nil {
		return "", err
	}

	return executeRequest(req)
}

func executeRequest(req *http.Request) (string, error) {
	cookie := viper.GetString("cookie")
	if cookie == "" {
		return "", utils.AOCCLIError("No cookie set. Please set the cookie using the --cookie flag or the config file")
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookie))
	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	cli.ToPrintf("Got status code %d", result.StatusCode).PrintDebug()
	if result.StatusCode != 200 {
		return "", utils.AOCCLIErrorf("Got status code %d", result.StatusCode)
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(result.Body)
	body := buffer.String()

	return string(body), nil
}

func GetDayPage(day int, year int, task int) (string, error) {
	htmlText, err := get(day, year, "")
	if err != nil {
		return "", err
	}
	wholeArticle := regexp.MustCompile("(?ms)<article(.*?)</article>").FindAllStringSubmatch(htmlText, -1)[task-1][1]
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

	// eliminate classname of body tag
	wholeArticle = regexp.MustCompile(" class=\"day-desc\">").ReplaceAllString(wholeArticle, "")

	return html.UnescapeString(wholeArticle), nil
}

func replaceTagRegex(text string, tagRegex string, replacement string) string {
	replacesRegex := regexp.MustCompile(">").ReplaceAllString(tagRegex, "[^>]*>")
	return regexp.MustCompile(replacesRegex).ReplaceAllString(text, replacement)
}

func GetSolveInput(day int, year int) (string, error) {
	result, err := get(day, year, "/input")
	if err != nil {
		return "", err
	}
	return strings.Trim(result, "\n"), nil
}

func GetTestInput(day int, year int) (string, error) {
	dayPage, err := GetResource("challenge1", day, year)
	if err != nil {
		return "", err
	}
	submatches := regexp.MustCompile("```([^`]*)```").FindStringSubmatch(dayPage)
	if submatches == nil || len(submatches) < 2 {
		cli.ToPrint("No test input found").PrintWarning()
		cli.ToPrint("An empty file was created. You can copy the test input into the file test.in").PrintDebug()
		return "", nil
	}
	return strings.Trim(submatches[1], "\n"), nil
}

func GetTestOutput(day int, year int, task int) (string, error) {
	dayPage, err := GetResource(fmt.Sprintf("challenge%d", task), day, year)
	if err != nil {
		return "", err
	}

	matches := regexp.MustCompile("[*][*]`[^`]*`[*][*]").FindStringSubmatch(dayPage)
	if matches == nil {
		cli.ToPrint("No test output found").PrintWarning()
		return "", nil
	}
	rawMatch := matches[len(matches)-1]
	return rawMatch[3 : len(rawMatch)-3], nil
}
