package utils

import "fmt"

type ExecutionDetails struct {
	Command          string
	Args             []string
	WorkingDirectory string
}

func (ed *ExecutionDetails) Arg(argument string) *ExecutionDetails {
	ed.Args = append(ed.Args, argument)
	return ed
}

func (ed *ExecutionDetails) Argf(argument string, a ...any) *ExecutionDetails {
	ed.Args = append(ed.Args, fmt.Sprintf(argument, a...))
	return ed
}

func (ed *ExecutionDetails) Dir(directory string) *ExecutionDetails {
	ed.WorkingDirectory = directory
	return ed
}

func (ed *ExecutionDetails) Dirf(directory string, a ...any) *ExecutionDetails {
	ed.WorkingDirectory = fmt.Sprintf(directory, a...)
	return ed
}

func ToExecute(command string) *ExecutionDetails {
	return &ExecutionDetails{
		Command: command,
		Args:    []string{},
	}
}
