package languages

import (
	"aoc-cli/utils"
	_ "embed"
)

type Python struct{}

//go:embed python/runner.py
var pythonRunnerFile string

//go:embed python/task.py
var pythonTaskFile string

func (p Python) GetSolveCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("python").Arg("-u").Arg("runner.py").Argf("%d", task)
}

func (p Python) GetTestCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("python").Arg("-u").Arg("runner.py").Argf("%d", task).Arg("test")
}

func (p Python) GetPreparationCommand(executionDirectory string, task int) []utils.ExecutionDetails {
	return []utils.ExecutionDetails{}
}

func (p Python) GetFilesToWrite() []utils.FileTemplate {
	runnerFile := utils.FileTemplate{Content: pythonRunnerFile, Filename: "runner.py"}
	taskFile := utils.FileTemplate{Content: pythonTaskFile, Filename: "task.py"}

	return []utils.FileTemplate{runnerFile, taskFile}
}
