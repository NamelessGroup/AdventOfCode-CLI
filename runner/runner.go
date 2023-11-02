package runner

import (
	"aoc-cli/runner/languages"
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/creack/pty"
)

type Language interface {
	GetSolveCommand(day int, task int) string
	GetTestCommand(day int, task int) string
	GetPreparationCommand(day int, task int) string
}

type RunResult struct {
	stdout            []string
	exitCode          int
	executionDuration time.Duration
}

func formatCommand(fullCommand string) (string, []string) {
	commandList := strings.Split(fullCommand, " ")
	if len(commandList) <= 0 {
		return "", []string{} // TODO: Throw error?
	}
	if len(commandList) <= 1 {
		return commandList[0], []string{}
	}
	return commandList[0], commandList[1:]
}

func runCommand(streamOutput bool, command string, args ...string) RunResult {
	cmd := exec.Command(command, args...)

	output := []string{}

	timeStart := time.Now()
	ptmx, err := pty.Start(cmd)

	if err != nil {
		log.Fatalf("cmd.Start() failed!")
	}

	scanner := bufio.NewScanner(ptmx)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		if streamOutput {
			fmt.Println(m)
		}
		output = append(output, m)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed!")
	}
	timeEnd := time.Now()
	timeTaken := timeEnd.Sub(timeStart)

	return RunResult{stdout: output, exitCode: cmd.ProcessState.ExitCode(), executionDuration: timeTaken}
}

func resolveLanguage(lang string) Language {
	if lang == "test" {
		return languages.Test{}
	}
	// TODO: Resolve language string to language object
	return nil
}

func prepareTask(day int, task int, lang Language) {
	rawCommand := lang.GetPreparationCommand(day, task)
	if rawCommand == "" {
		return
	}
	fmt.Println("Preparing Day", day, "Task", task)
	command, args := formatCommand(rawCommand)
	result := runCommand(false, command, args...)

	if result.exitCode == 0 {
		fmt.Println("Prepared sucessfully")
	} else {
		fmt.Println("Preparation failed with exit code", result.exitCode)
	}
}

func runTask(day int, task int, rawCommand string) {
	fmt.Println("Running day", day, "Task", task)
	command, args := formatCommand(rawCommand)
	result := runCommand(true, command, args...)

	if result.exitCode == 0 {
		fmt.Println("Task", task, "finished successfully after", result.executionDuration.Truncate(10000))
	} else {
		fmt.Println("Task", task, "failed execution after", result.executionDuration.Truncate(100000), " with exit code", result.exitCode)
	}
}

func SolveDay(day int, task int, lang string) {
	languageObject := resolveLanguage(lang)

	if languageObject == nil {
		return // TODO: Error throwing?
	}

	prepareTask(day, task, languageObject)
	rawRunCommand := languageObject.GetSolveCommand(day, task)
	runTask(day, task, rawRunCommand)
}

func TestDay(day int, task int, lang string) {
	languageObject := resolveLanguage(lang)

	if languageObject == nil {
		return // TODO: Error throwing?
	}

	prepareTask(day, task, languageObject)
	rawRunCommand := languageObject.GetTestCommand(day, task)
	runTask(day, task, rawRunCommand)
}
