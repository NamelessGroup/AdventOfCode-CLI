package cli

import (
	"os"
	"github.com/fatih/color"
)

type Format struct {
	underline bool
	bold bool
	italic bool
}

func PrintError(message string) {
	Print(message, color.FgRed, Format{false, false, false}, true)
	os.Exit(1)
}

func PrintLog(message string, breakAtEnd bool) {
	Print(message, color.FgWhite, Format{false, false, false}, breakAtEnd)
}

func PrintDebug(message string) {
	Print(message, color.FgCyan, Format{false, false, true}, true)
}

func PrintSuccess(message string) {
	Print(message, color.FgGreen, Format{false, true, false}, true)
}

func PrintWarning(message string) {
	Print(message, color.FgYellow, Format{false, false, false}, true)
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
	c.Printf("\r%s", message)
	if (breakAtEnd) {
		c.Println()
	}
}