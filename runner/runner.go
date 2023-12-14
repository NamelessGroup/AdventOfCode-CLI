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

func runCommand(streamOutput bool, toRun utils.ExecutionDetails, workingDirectory string, s *cli.Spinner) RunResult {
	cmd := exec.Command(toRun.Command, toRun.Args...)

	output := []string{}

	cli.ToPrintf("Running command %s with args %s", toRun.Command, toRun.Args).PrintDebug()

	if toRun.WorkingDirectory == "" {
		cmd.Dir = workingDirectory
	} else {
		cmd.Dir = toRun.WorkingDirectory
	}
	cli.ToPrintf("Setting working directory to %s", cmd.Dir).PrintDebug()

	timeStart := time.Now()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cli.PrintFromError(err).PrintDebug()
		cli.ToPrint("Error getting stdout pipe!").PrintError()
	}

	stderr, err := cmd.StderrPipe()

	if err != nil {
		cli.PrintFromError(err).PrintDebug()
		cli.ToPrint("Error getting stderr pipe!").PrintError()
	}

	err = cmd.Start()

	if err != nil {
		cli.PrintFromError(err).PrintDebug()
		cli.ToPrint("Error starting command!").PrintError()
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		if streamOutput {
			cli.ToPrint(m).Color(color.FgCyan).Italic().Print()
			if s != nil {
				s.Reprint()
			}
		}
		output = append(output, m)
	}

	errScanner := bufio.NewScanner(stderr)

	for errScanner.Scan() {
		m := errScanner.Text()
		if streamOutput {
			cli.ToPrint(m).Color(color.FgCyan).Italic().Print()
			if s != nil {
				s.Reprint()
			}
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
	cli.ToPrintf("Preparing day %d task %d", day, task).PrintLog()

	preparedSuccessfully := true

	for _, executionDetails := range rawCommand {
		result := runCommand(false, executionDetails, executionDirectory, nil)

		if result.exitCode != 0 {
			cli.ToPrintf("Preparation failed with exit code %d", result.exitCode).PrintError()
			preparedSuccessfully = false
			break
		}
	}

	if preparedSuccessfully {
		cli.ToPrint("Successfully prepared!").PrintSuccess()
	}
}

func runTask(day int, task int, executionDetails utils.ExecutionDetails, executionDirectory string) []string {
	s := cli.Spinner{}
	s.Run(fmt.Sprintf("Running day %d task %d", day, task))
	result := runCommand(true, executionDetails, executionDirectory, &s)
	s.Stop()
	if result.exitCode == 0 {
		cli.ToPrintf("Task %d finished successfully after %s", task, result.executionDuration.Truncate(10000)).PrintSuccess()
	} else {
		cli.ToPrintf("Task %d failed execution after %s with exit code %d", task, result.executionDuration.Truncate(10000), result.exitCode).PrintError()
	}

	if len(result.stdout) == 0 {
		cli.ToPrint("Your code didn't output anything!").PrintWarning()
		return []string{""}
	} else {
		return result.stdout
	}
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
