package aocweb

import (
	"aoc-cli/cli/flags"
	cli "aoc-cli/output"
	"aoc-cli/utils"
	"bytes"
	"fmt"
	"regexp"

	//"io"
	"net/http"

	"github.com/spf13/viper"
)

func postAnswer(day int, year int, task int, solution string) (string, error) {
	cli.ToPrintf("Submitting solution for day %d year %d", day, year).PrintDebug()
	reqBody := []byte(fmt.Sprintf("level=%d&answer=%s", task, solution))
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day), bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	cookie := viper.GetString(flags.Flags["cookie"].ViperKey)
	if cookie == "" {
		return "", utils.AOCCLIError("No cookie set. Please set the cookie using the --cookie flag or the config file")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookie))

	return executeRequest(req)
}

func Submit(day int, year int, task int, solution string) error {
	result, err := postAnswer(day, year, task, solution)
	if err != nil {
		return utils.AOCCLIError("Error submitting answer!").DebugInfo("submitter", err.Error())
	}

	wholeArticle := regexp.MustCompile("(?ms)<article(.*)</article>").FindString(result)
	wholeArticle = replaceTagRegex(wholeArticle, "</?article>", "")
	wholeArticle = replaceTagRegex(wholeArticle, "</?p>", "")

	if (regexp.MustCompile("That's the right answer")).MatchString(wholeArticle) {
		return nil
	}
	if (regexp.MustCompile("You don't seem to be solving the right level")).MatchString(wholeArticle) {
		return utils.AOCCLIErrorf("Already solved task %d for day %d of %d", task, day, year)
	}
	if (regexp.MustCompile("That's not the right answer")).MatchString(wholeArticle) {
		return utils.AOCCLIErrorf("Wrong answer submitted.\nYour answer: %s", solution)
	}
	if (regexp.MustCompile("You gave an answer too recently")).MatchString(wholeArticle) {
		timeToWait := regexp.MustCompile("You have (.*) left to wait").FindStringSubmatch(wholeArticle)
		formattedTimeToWait := "an unknown timespan"
		if timeToWait == nil {
			timeToWait = regexp.MustCompile("You have (.*) left to wait").FindStringSubmatch(wholeArticle)
			if timeToWait != nil {
				formattedTimeToWait = timeToWait[1]
			}
		} else {
			formattedTimeToWait = timeToWait[1]
		}
		return utils.AOCCLIErrorf("You gave an answer too recently. You have to wait for %s to submit again", formattedTimeToWait)
	}
	return utils.AOCCLIErrorf("Unknown error submitting answer: %s", replaceTagRegex(wholeArticle, "<>", ""))
}
