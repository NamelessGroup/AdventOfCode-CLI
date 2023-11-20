package languages

import (
	"aoc-cli/utils"
	_ "embed"
)

type Java struct{}

//go:embed java/Runner.java
var javaRunnerFile string

//go:embed java/Tasks.java
var javaTaskFile string

func (j Java) GetSolveCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("java").Arg("Runner").Argf("%d", task).Arg(".").Dir(executionDirectory)
}

func (j Java) GetTestCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("java").Arg("Runner").Argf("%d", task).Arg(".").Arg("test").Dir(executionDirectory)
}

func (j Java) GetPreparationCommand(executionDirectory string, task int) []utils.ExecutionDetails {
	return []utils.ExecutionDetails{
		*utils.ToExecute("javac").Arg("Runner.java").Arg("Tasks.java").Dir(executionDirectory),
	}
}

func (j Java) GetFilesToWrite() []utils.FileTemplate {
	runnerFile := utils.FileTemplate{Content: javaRunnerFile, Filename: "Runner.java"}
	taskFile := utils.FileTemplate{Content: javaTaskFile, Filename: "Tasks.java"}

	return []utils.FileTemplate{runnerFile, taskFile}
}
