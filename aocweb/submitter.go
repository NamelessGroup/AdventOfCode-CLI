package aocweb

import (
	cli "aoc-cli/output"
	"bytes"
	"errors"
	"fmt"
	"regexp"

	//"io"
	"net/http"

	"github.com/spf13/viper"
)

func postAnswer(day int, year int, task int, solution string) (string, error) {
	cli.PrintDebugFmt("Submitting solution for day %d year %d", day, year)
	reqBody := []byte(fmt.Sprintf("level=%d&answer=%s", task, solution))
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day), bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	cookie := viper.GetString("cookie")
	if cookie == "" {
		return "", fmt.Errorf("No cookie set. Please set the cookie using the --cookie flag or the config file")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookie))

	return executeRequest(req)
}

func Submit(day int, year int, task int, solution string) error {
	result, err := postAnswer(day, year, task, solution)
	if err != nil {
		cli.PrintDebug(err.Error())
		return errors.New("Error submitting answer!")
	}

	wholeArticle := regexp.MustCompile("(?ms)<article(.*)</article>").FindString(result)
	wholeArticle = replaceTagRegex(wholeArticle, "</?article>", "")
	wholeArticle = replaceTagRegex(wholeArticle, "</?p>", "")

	if (regexp.MustCompile("That's the right answer")).MatchString(wholeArticle) {
		return nil
	}
	if (regexp.MustCompile("You don't seem to be solving the right level")).MatchString(wholeArticle) {
		return fmt.Errorf("Already solved task %d for day %d of %d", task, day, year)
	}
	if (regexp.MustCompile("That's not the right answer")).MatchString(wholeArticle) {
		return fmt.Errorf("Wrong answer submitted.\nYour answer: %s", solution)
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
		return fmt.Errorf("You gave an answer too recently. You have to wait for %s to submit again", formattedTimeToWait)
	}
	return fmt.Errorf("Unknown error submitting answer: %s", replaceTagRegex(wholeArticle, "<>", ""))
}
