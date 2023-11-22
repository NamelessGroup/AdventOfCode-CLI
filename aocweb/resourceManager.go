package aocweb

import (
	cli "aoc-cli/output"
	"aoc-cli/utils"
	"fmt"
	"os"
)

type Resource struct {
	fileName string
	getter   func(day int, year int) (string, error)
}

var resources = make(map[string]Resource)

func init() {
	resources["challenge1"] = Resource{
		fileName: "challenge1.md",
		getter:   (func(day int, year int) (string, error) {
			return GetDayPage(day, year, 1)
		}),
	}
	resources["challenge2"] = Resource{
		fileName: "challenge2.md",
		getter:   (func(day int, year int) (string, error) {
			return GetDayPage(day, year, 2)
		}),
	}
	resources["solveInput"] = Resource{
		fileName: "solve.in",
		getter:   GetSolveInput,
	}
	resources["testInput"] = Resource{
		fileName: "test.in",
		getter:   GetTestInput,
	}
	resources["testOutput1"] = Resource{
		fileName: "testOne.out",
		getter: (func(day int, year int) (string, error) {
			return GetTestOutput(day, year, 1)
		}),
	}
	resources["testOutput2"] = Resource{
		fileName: "testTwo.out",
		getter: (func(day int, year int) (string, error) {
			return GetTestOutput(day, year, 2)
		}),
	}
}

func GetResource(name string, day int, year int) (string, error) {
	cli.PrintDebugFmt("Getting resource %s for day %d year %d", name, day, year)
	resource, foundResource := resources[name]
	if !foundResource {
		return "", utils.AOCCLIErrorf("Resource %s not registered", name)
	}

	cli.PrintDebugFmt("Looking for file %d/%d/%s", year, day, resource.fileName)
	dir := utils.GetChallengeDirectory(year, day)
	fileContent, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, resource.fileName))
	fileContentString := string(fileContent)
	if err == nil && fileContentString == "" {
		cli.PrintWarningFmt("Local file of resource %s found, but was empty", fmt.Sprintf("%d/%d/%s", year, day, resource.fileName))
	}
	if err == nil && fileContentString != "" {
		cli.PrintDebugFmt("Local file of resource %s found", name)
		cli.PrintDebug(fileContentString)
		return fileContentString, nil
	}
	cli.PrintDebugFmt("Local file of resource %s not found", name)
	resourceContent, err := resource.getter(day, year)
	if err != nil {
		return "", err
	}

	cli.PrintDebugFmt("Found Resource %s:", name)
	cli.PrintDebug(resourceContent)
	cli.PrintDebugFmt("Saving resource %s", name)

	resource.saveToFile(year, day, resourceContent)
	return resourceContent, nil
}

func (resource Resource) saveToFile(year int, day int, content string) {
	dir := utils.GetChallengeDirectory(year, day)
	err := os.WriteFile(fmt.Sprintf("%s/%s", dir, resource.fileName), []byte(content), 0644)
	if err != nil {
		cli.PrintDebug("Could not save resource to file:")
		cli.PrintDebugError(err)
	}
}
