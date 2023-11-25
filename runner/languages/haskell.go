package languages

import (
	"aoc-cli/utils"
	_ "embed"
)

type Haskell struct{}

//go:embed haskell/runner.hs
var haskellRunnerFile string

//go:embed haskell/task.hs
var haskellTaskFile string

func (h Haskell) GetSolveCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("ghc").Arg("--run").Arg("runner.hs").Arg("--").Argf("%d", task).Arg("solve")
}

func (h Haskell) GetTestCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("ghc").Arg("--run").Arg("runner.hs").Arg("--").Argf("%d", task).Arg("test")
}

func (h Haskell) GetPreparationCommand(executionDirectory string, task int) []utils.ExecutionDetails {
	return []utils.ExecutionDetails{}
}

func (h Haskell) GetFilesToWrite() []utils.FileTemplate {
	runnerFile := utils.FileTemplate{Content: haskellRunnerFile, Filename: "runner.hs"}
	taskFile := utils.FileTemplate{Content: haskellTaskFile, Filename: "task.hs"}

	return []utils.FileTemplate{runnerFile, taskFile}
}

func (h Haskell) GetLanguageSpecificConfigKeys() map[string]utils.FlagMetadata {
	return map[string]utils.FlagMetadata{}
}
