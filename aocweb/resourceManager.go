package aocweb

import (
	"fmt"
	"os"
	"aoc-cli/output"
)

type Resource struct {
	fileName string
	getter   func(day int, year int) (string, error)
}

var resources = make(map[string]Resource)

func init() {
	resources["challenge"] = Resource{
		fileName: "challenge.md",
		getter: GetDayPage,
	}
	resources["solveInput"] = Resource{
		fileName: "solve.in",
		getter: GetSolveInput,
	}
	resources["testInput"] = Resource{
		fileName: "test.in",
		getter: GetTestInput,
	}
}

func GetResource(name string, day int, year int) (string, error) {
	cli.PrintDebugFmt("Getting resource %s for day %d year %d", name, day, year)
	resource, foundResource := resources[name]
	if !foundResource {
		return "", fmt.Errorf("Resource %s not registered", name)
	}

	cli.PrintDebugFmt("Looking for file %d/%d/%s", year, day, resource.fileName)
	fileContent, err := os.ReadFile(fmt.Sprintf("%d/%d/%s", year, day, resource.fileName))
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
	err := os.WriteFile(fmt.Sprintf("%d/%d/%s", year, day, resource.fileName), []byte(content), 0644)
	if err != nil {
		cli.PrintDebug("Could not save resource to file:")
		cli.PrintDebug(err.Error())
	}
}