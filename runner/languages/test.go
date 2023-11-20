package languages

import (
	"aoc-cli/utils"
)

type Test struct{}

func (t Test) GetSolveCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("python3").Arg("./test.py")
}

func (t Test) GetTestCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("echo").Argf("This is the test command for path %s task %d", executionDirectory, task)
}

func (t Test) GetPreparationCommand(executionDirectory string, task int) []utils.ExecutionDetails {
	return []utils.ExecutionDetails{
		*utils.ToExecute("echo").Arg("This is never seen"),
	}
}

func (t Test) GetFilesToWrite() []utils.FileTemplate {
	return []utils.FileTemplate{}
}
