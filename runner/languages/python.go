package languages

import (
	"aoc-cli/utils"
	_ "embed"
	"fmt"
)

type Python struct{}

//go:embed python/runner.py
var runner string

//go:embed python/task.py
var task string

func (p Python) GetSolveCommand(executionDirectory string, task int) string {
	return fmt.Sprintf("python3 %srunner.py %d", executionDirectory, task)
}

func (p Python) GetTestCommand(executionDirectory string, task int) string {
	return fmt.Sprintf("python3 %srunner.py %d test", executionDirectory, task)
}

func (p Python) GetPreparationCommand(executionDirectory string, task int) []string {
	return []string{}
}

func (p Python) GetFilesToWrite() []utils.FileTemplate {
	runnerFile := utils.FileTemplate{Content: runner, Filename: "runner.py"}
	taskFile := utils.FileTemplate{Content: task, Filename: "task.py"}

	return []utils.FileTemplate{runnerFile, taskFile}
}
