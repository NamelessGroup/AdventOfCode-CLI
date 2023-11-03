package aocweb

import (
	cli "aoc-cli/output"
	"bytes"
	"errors"
	"fmt"

	"net/http"
	"regexp"
)

func get(day int, year int, path string) (string, error) {
	cli.PrintDebugFmt("Getting %s for day %d year %d from the website", path, day, year)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d%s", year, day, path), nil)
	if err != nil {
		return "", err
	}
	// req.Header.Set("Cookie", fmt.Sprintf("session=%s", COOKIE))
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
	return get(day, year, "")
}

func GetSolveInput(day int, year int) (string, error) {
	return get(day, year, "/input")
}

func GetTestInput(day int, year int) (string, error) {
	dayPage, err := GetResource("dayPage", day, year)
	if err != nil {
		return "", err
	}
	openMatch, err := firstInstance(dayPage, "<code")
	closeMatch, err := firstInstance(dayPage, "</code>")
	if err != nil {
		return "", errors.New("Could not find test input")
	}
	codeBlock := dayPage[openMatch:closeMatch]
	testInput, err := firstInstance(codeBlock, ">")
	if err != nil {
		return "", errors.New("Could not find test input")
	}
	return codeBlock[testInput+1:], nil
}

func firstInstance(text string, regEx string) (int, error) {
	re := regexp.MustCompile(regEx)
	matches := re.FindAllStringIndex(text, -1)
	if len(matches) < 1 {
		return -1, fmt.Errorf("No match found")
	}
	return matches[0][0], nil
}
