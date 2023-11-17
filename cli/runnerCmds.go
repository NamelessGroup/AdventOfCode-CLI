package cli

import (
	"aoc-cli/utils"
	"strconv"
)

func getTask(args []string) (int, error) {
	if len(args) == 0 {
		return 1, nil
	}

	task, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, err
	}

	if task != 1 && task != 2 {
		return 0, utils.AOCCLIError("Task must be 1 or 2").DebugInfof("runnerCmds", "Supplied task: %d", task)
	}
	return task, nil
}
