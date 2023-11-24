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
		getter: (func(day int, year int) (string, error) {
			return GetDayPage(day, year, 1)
		}),
	}
	resources["challenge2"] = Resource{
		fileName: "challenge2.md",
		getter: (func(day int, year int) (string, error) {
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
	cli.ToPrintf("Getting resource %s for day %d year %d", name, day, year).PrintDebug()
	resource, foundResource := resources[name]
	if !foundResource {
		return "", utils.AOCCLIErrorf("Resource %s not registered", name)
	}

	cli.ToPrintf("Looking for file %d/%d/%s", year, day, resource.fileName).PrintDebug()
	dir := utils.GetChallengeDirectory(year, day)
	fileContent, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, resource.fileName))
	fileContentString := string(fileContent)
	if err == nil && fileContentString == "" {
		cli.ToPrintf("Local file of resource %s found, but was empty", fmt.Sprintf("%d/%d/%s", year, day, resource.fileName)).PrintWarning()
	}
	if err == nil && fileContentString != "" {
		cli.ToPrintf("Local file of resource %s found", name).PrintDebug()
		cli.ToPrint(fileContentString).PrintDebug()
		return fileContentString, nil
	}
	cli.ToPrintf("Local file of resource %s not found", name).PrintDebug()
	resourceContent, err := resource.getter(day, year)
	if err != nil {
		return "", err
	}

	cli.ToPrintf("Found Resource %s:", name).PrintDebug()
	cli.ToPrint(resourceContent).PrintDebug()
	cli.ToPrintf("Saving resource %s", name).PrintDebug()

	resource.saveToFile(year, day, resourceContent)
	return resourceContent, nil
}

func (resource Resource) saveToFile(year int, day int, content string) {
	dir := utils.GetChallengeDirectory(year, day)
	err := os.WriteFile(fmt.Sprintf("%s/%s", dir, resource.fileName), []byte(content), 0644)
	if err != nil {
		cli.ToPrint("Could not save resource to file:").PrintDebug()
		cli.PrintFromError(err).PrintDebug()
	}
}
