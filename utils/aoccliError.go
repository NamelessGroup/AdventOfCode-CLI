package utils

import "fmt"

type AOC_CLIError struct {
	Message   string
	module    string
	debugInfo string
}

func (pe *AOC_CLIError) Error() string {
	return pe.Message
}

func (pe *AOC_CLIError) DebugInfo(module string, debugInfo string) *AOC_CLIError {
	pe.debugInfo = debugInfo
	pe.module = module
	return pe
}

func (pe *AOC_CLIError) DebugInfof(module string, debugInfo string, a ...any) *AOC_CLIError {
	pe.debugInfo = fmt.Sprintf(debugInfo, a...)
	pe.module = module
	return pe
}

func (pe *AOC_CLIError) DebugError() string {
	return fmt.Sprintf("%s: %s", pe.module, pe.debugInfo)
}

func (pe *AOC_CLIError) HasDebugInfo() bool {
	return pe.debugInfo != ""
}

func AOCCLIError(message string) *AOC_CLIError {
	return &AOC_CLIError{
		Message: message,
	}
}

func AOCCLIErrorf(message string, a ...any) *AOC_CLIError {
	return &AOC_CLIError{
		Message: fmt.Sprintf(message, a...),
	}
}
