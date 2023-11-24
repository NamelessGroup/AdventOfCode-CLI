package runner

import (
	cli "aoc-cli/output"
	utils "aoc-cli/utils"
	"bufio"
	"fmt"

	"os/exec"
	"time"

	"github.com/fatih/color"
)

type RunResult struct {
	stdout            []string
	exitCode          int
	executionDuration time.Duration
}

func runCommand(streamOutput bool, toRun utils.ExecutionDetails, workingDirectory string) RunResult {
	cmd := exec.Command(toRun.Command, toRun.Args...)

	output := []string{}

	cli.PrintDebugFmt("Running command %s with args %s", toRun.Command, toRun.Args)

	if toRun.WorkingDirectory == "" {
		cmd.Dir = workingDirectory
	} else {
		cmd.Dir = toRun.WorkingDirectory
	}
	cli.PrintDebugFmt("Setting working directory to %s", cmd.Dir)

	timeStart := time.Now()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cli.PrintDebugError(err)
		cli.PrintErrorString("Error getting stdout pipe!")
	}

	stderr, err := cmd.StderrPipe()

	if err != nil {
		cli.PrintDebugError(err)
		cli.PrintErrorString("Error getting stderr pipe!")
	}

	err = cmd.Start()

	if err != nil {
		cli.PrintDebugError(err)
		cli.PrintErrorString("Error starting command!")
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		if streamOutput {
			cli.PrintRaw(m, cli.Format{
				Color:   color.FgCyan,
				NewLine: true,
			})
		}
		output = append(output, m)
	}

	errScanner := bufio.NewScanner(stderr)

	for errScanner.Scan() {
		m := errScanner.Text()
		if streamOutput {
			cli.PrintRaw(m, cli.Format{
				Color:   color.FgCyan,
				NewLine: true,
			})
		}
		output = append(output, m)
	}

	cmd.Wait()

	timeEnd := time.Now()
	timeTaken := timeEnd.Sub(timeStart)

	return RunResult{stdout: output, exitCode: cmd.ProcessState.ExitCode(), executionDuration: timeTaken}
}

func prepareTask(year int, day int, task int, lang Language) {
	executionDirectory := utils.GetChallengeDirectory(year, day)
	rawCommand := lang.GetPreparationCommand(executionDirectory, task)
	if len(rawCommand) == 0 {
		return
	}
	cli.PrintLogFmt("Preparing day %d task %d", day, task)

	preparedSuccessfully := true

	for _, executionDetails := range rawCommand {
		result := runCommand(false, executionDetails, executionDirectory)

		if result.exitCode != 0 {
			cli.PrintErrorFmt("Preparation failed with exit code %d", result.exitCode)
			preparedSuccessfully = false
			break
		}
	}

	if preparedSuccessfully {
		cli.PrintSuccess("Successfully prepared!")
	}
}

func runTask(day int, task int, executionDetails utils.ExecutionDetails, executionDirectory string) []string {
	s := cli.Spinner{}
	s.Run(fmt.Sprintf("Running day %d task %d", day, task))
	result := runCommand(true, executionDetails, executionDirectory)
	s.Stop()
	if result.exitCode == 0 {
		cli.PrintSuccessFmt("Task %d finished successfully after %s", task, result.executionDuration.Truncate(10000))
	} else {
		cli.PrintErrorFmt("Task %d failed execution after %s with exit code %d", task, result.executionDuration.Truncate(10000), result.exitCode)
	}
	return result.stdout
}

func SolveDay(year int, day int, task int, languageObject Language) []string {
	prepareTask(year, day, task, languageObject)
	executionDirectory := utils.GetChallengeDirectory(year, day)
	rawRunCommand := languageObject.GetSolveCommand(executionDirectory, task)
	return runTask(day, task, rawRunCommand, executionDirectory)
}

func TestDay(year int, day int, task int, languageObject Language) []string {
	prepareTask(year, day, task, languageObject)
	executionDirectory := utils.GetChallengeDirectory(year, day)
	rawRunCommand := languageObject.GetTestCommand(executionDirectory, task)
	return runTask(day, task, rawRunCommand, executionDirectory)
}
