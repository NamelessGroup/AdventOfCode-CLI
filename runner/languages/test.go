package languages

import (
	"aoc-cli/utils"
	"fmt"
)

type Test struct{}

func (t Test) GetSolveCommand(executionDirectory string, task int) string {
	return "python3 ./test.py"
}

func (t Test) GetTestCommand(executionDirectory string, task int) string {
	return fmt.Sprintf("echo This is the test command for path %s task %d", executionDirectory, task)
}

func (t Test) GetPreparationCommand(executionDirectory string, task int) []string {
	return []string{"echo This is never seen"}
}

func (t Test) GetFilesToWrite() []utils.FileTemplate {
	return []utils.FileTemplate{}
}
