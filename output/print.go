package cli

import (
	"aoc-cli/utils"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	ColorLog     = color.FgWhite
	ColorWarning = color.FgYellow
	ColorError   = color.FgRed
	ColorDebug   = color.FgMagenta
	ColorSuccess = color.FgGreen
)

type Format struct {
	Underline bool
	Bold      bool
	Italic    bool
	Color     color.Attribute
	NewLine   bool
	Append    bool
}

var PrintDebugMessages = false
var DisableEmojis = false

func PrintErrorString(message string) {
	if !DisableEmojis {
		message = "❕ " + message
	}
	PrintRaw(message, Format{
		Color:   ColorError,
		NewLine: true,
	})
	os.Exit(1)
}

func PrintErrorFmt(message string, a ...any) {
	PrintErrorString(fmt.Sprintf(message, a...))
}

func PrintError(err error) {
	var aocCliError *utils.AOC_CLIError
	if errors.As(err, &aocCliError) {
		PrintDebug(aocCliError.DebugError())
	}
	PrintErrorString(err.Error())
}

func PrintLog(message string, newline bool) {
	if !DisableEmojis {
		message = "ℹ " + message
	}
	PrintRaw(message, Format{
		Color:   ColorLog,
		NewLine: newline,
	})
}

func PrintLogFmt(message string, a ...any) {
	PrintLog(fmt.Sprintf(message, a...), true)
}

func PrintDebug(message string) {
	if PrintDebugMessages {
		if !DisableEmojis {
			message = "🐞 " + message
		}
		PrintRaw(message, Format{
			Color:   ColorDebug,
			NewLine: true,
		})
	}
}

func PrintDebugError(err error) {
	var aocCliError *utils.AOC_CLIError
	if errors.As(err, &aocCliError) {
		PrintDebug(aocCliError.DebugError())
	}
	PrintDebug(err.Error())
}

func PrintDebugFmt(message string, a ...any) {
	PrintDebug(fmt.Sprintf(message, a...))
}

func PrintSuccess(message string) {
	if !DisableEmojis {
		message = "✔️ " + message
	}
	PrintRaw(message, Format{
		Color:   ColorSuccess,
		NewLine: true,
	})
}

func PrintSuccessFmt(message string, a ...any) {
	PrintSuccess(fmt.Sprintf(message, a...))
}

func PrintWarning(message string) {
	if !DisableEmojis {
		message = "⚠️ " + message
	}
	PrintRaw(message, Format{
		Color:   ColorWarning,
		NewLine: true,
	})
}

func PrintWarningFmt(message string, a ...any) {
	PrintWarning(fmt.Sprintf(message, a...))
}

func PrintRaw(message string, format Format) {
	c := color.New()
	if format.Color != 0 {
		c.Add(format.Color)
	}
	if format.Bold {
		c.Add(color.Bold)
	}
	if format.Italic {
		c.Add(color.Italic)
	}
	if format.Underline {
		c.Add(color.Underline)
	}

	if !format.Append {
		c.Print("\r\033[K")
	}
	c.Print(message)
	if format.NewLine {
		c.Println()
	}
}
