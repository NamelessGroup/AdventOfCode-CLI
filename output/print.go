package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Format struct {
	underline bool
	bold      bool
	italic    bool
}

var PrintDebugMessages = false

func PrintError(message string) {
	Print(message, color.FgRed, Format{false, false, false}, true)
	os.Exit(1)
}

func PrintErrorFmt(message string, a ...any) {
	PrintError(fmt.Sprintf(message, a...))
}

func PrintLog(message string, breakAtEnd bool) {
	Print(message, color.FgWhite, Format{false, false, false}, breakAtEnd)
}

func PrintLogFmt(message string, a ...any) {
	PrintLog(fmt.Sprintf(message, a...), true)
}

func PrintDebug(message string) {
	if PrintDebugMessages {
		Print(message, color.FgMagenta, Format{false, false, true}, true)
	}
}

func PrintDebugFmt(message string, a ...any) {
	PrintDebug(fmt.Sprintf(message, a...))
}

func PrintSuccess(message string) {
	Print(message, color.FgGreen, Format{false, true, false}, true)
}

func PrintSuccessFmt(message string, a ...any) {
	PrintSuccess(fmt.Sprintf(message, a...))
}

func PrintWarning(message string) {
	Print(message, color.FgYellow, Format{false, false, false}, true)
}

func PrintWarningFmt(message string, a ...any) {
	PrintWarning(fmt.Sprintf(message, a...))
}

func Print(message string, col color.Attribute, format Format, breakAtEnd bool) {
	c := color.New(col)
	if format.underline {
		c.Add(color.Underline)
	}
	if format.bold {
		c.Add(color.Bold)
	}
	if format.italic {
		c.Add(color.Italic)
	}
	c.Printf("\r\033[K%s", message)
	if breakAtEnd {
		c.Println()
	}
}
