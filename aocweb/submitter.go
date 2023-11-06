package aocweb

import (
	"aoc-cli/output"
	"bytes"
	"errors"
	"fmt"
	"regexp"

	//"io"
	"net/http"
)

func postAnswer(day int, year int, task int, solution string) (string, error) {
	cli.PrintLog("Test", true)
	cli.PrintDebugFmt("Submitting solution for day %d year %d", day, year)
	reqBody := []byte(fmt.Sprintf("level=%d&answer=%s", task, solution))
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day), bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", "53616c7465645f5fde78c78e110a80e9cb826caf1faa75e964845a07beef93c1bc2d2174ef7ef645c665df8af64a9dd7961805ac0957109dcc3a437994b39a11"))
	
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
		timeToWait := regexp.MustCompile("You have to wait (.*) to submit again").FindStringSubmatch(wholeArticle)[1]
		return fmt.Errorf("You gave an answer too recently. You have to wait for %s to submit again", timeToWait)
	}
	return fmt.Errorf("Unknown error submitting answer: %s", replaceTagRegex(wholeArticle, "<>", ""))
}

