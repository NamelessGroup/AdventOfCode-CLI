package languages

import (
	"aoc-cli/utils"
	_ "embed"
)

type Rust struct{}

//go:embed rust/runner.rs
var rustRunnerFile string

//go:embed rust/task.rs
var rustTaskFile string

func (r Rust) GetSolveCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("./main").Argf("%d", task)
}

func (r Rust) GetTestCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("./main").Argf("%d", task).Arg("test")
}

func (r Rust) GetPreparationCommand(executionDirectory string, task int) []utils.ExecutionDetails {
	return []utils.ExecutionDetails{
		*utils.ToExecute("rustc").Arg("-O").Arg("main.rs"),
	}
}

func (r Rust) GetFilesToWrite() []utils.FileTemplate {
	runnerFile := utils.FileTemplate{Content: rustRunnerFile, Filename: "/main.rs"}
	taskFile := utils.FileTemplate{Content: rustTaskFile, Filename: "/task.rs"}

	return []utils.FileTemplate{runnerFile, taskFile}
}

func (r Rust) GetLanguageSpecificConfigKeys() map[string]utils.LanguageSpecificOption {
	return map[string]utils.LanguageSpecificOption{}
}
