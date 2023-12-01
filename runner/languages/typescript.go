package languages

import (
	"aoc-cli/utils"
	_ "embed"
)

type TypeScript struct{}

//go:embed typescript/runner.ts
var tsRunnerFile string

//go:embed typescript/task.ts
var tsTaskFile string

func (ts TypeScript) GetSolveCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("npx").Arg("ts-node").Arg("runner.ts").Argf("%d", task)
}

func (ts TypeScript) GetTestCommand(executionDirectory string, task int) utils.ExecutionDetails {
	return *utils.ToExecute("npx").Arg("ts-node").Arg("runner.ts").Argf("%d", task).Arg("test")
}

func (ts TypeScript) GetPreparationCommand(executionDirectory string, task int) []utils.ExecutionDetails {
	return []utils.ExecutionDetails{}
}

func (ts TypeScript) GetFilesToWrite() []utils.FileTemplate {
	runnerFile := utils.FileTemplate{Content: tsRunnerFile, Filename: "runner.ts"}
	taskFile := utils.FileTemplate{Content: tsTaskFile, Filename: "task.ts"}

	return []utils.FileTemplate{runnerFile, taskFile}
}

func (ts TypeScript) GetLanguageSpecificConfigKeys() map[string]utils.LanguageSpecificOption {
	return map[string]utils.LanguageSpecificOption{}
}
