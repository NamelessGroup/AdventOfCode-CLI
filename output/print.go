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

type Printable struct {
	underline bool
	bold      bool
	italic    bool
	color     color.Attribute
	newline   bool
	append    bool
	text      string
}

var PrintDebugMessages = false
var DisableEmojis = false

func ToPrint(message string) *Printable {
	return &Printable{
		underline: false,
		bold:      false,
		italic:    false,
		color:     color.FgWhite,
		newline:   true,
		append:    false,
		text:      message,
	}
}

func ToPrintf(message string, a ...any) *Printable {
	return ToPrint(fmt.Sprintf(message, a...))
}

func PrintFromError(err error) *Printable {
	var aocCliError *utils.AOC_CLIError
	if errors.As(err, &aocCliError) {
		ToPrint(aocCliError.DebugError()).PrintDebug()
	}
	return ToPrint(err.Error())
}

func (p *Printable) Underlined() *Printable {
	return p.SetUnderlined(true)
}

func (p *Printable) SetUnderlined(underline bool) *Printable {
	p.underline = underline
	return p
}

func (p *Printable) Bold() *Printable {
	return p.SetBold(true)
}

func (p *Printable) SetBold(bold bool) *Printable {
	p.bold = bold
	return p
}

func (p *Printable) Italic() *Printable {
	return p.SetItalic(true)
}

func (p *Printable) SetItalic(italic bool) *Printable {
	p.italic = italic
	return p
}

func (p *Printable) Color(col color.Attribute) *Printable {
	p.color = col
	return p
}

func (p *Printable) NewLine(newline bool) *Printable {
	p.newline = newline
	return p
}

func (p *Printable) Append(append bool) *Printable {
	p.append = append
	return p
}

func (p *Printable) Print() {
	c := color.New()
	c.Add(p.color)
	if p.bold {
		c.Add(color.Bold)
	}
	if p.italic {
		c.Add(color.Italic)
	}
	if p.underline {
		c.Add(color.Underline)
	}

	if !p.append {
		c.Print("\r\033[K")
	}
	c.Print(p.text)
	if p.newline {
		c.Println()
	}
}

func (p *Printable) PrintLog() {
	if !DisableEmojis {
		p.text = "ℹ " + p.text
	}
	p.color = ColorLog
	p.Print()
}

func (p *Printable) PrintSuccess() {
	if !DisableEmojis {
		p.text = "✔️ " + p.text
	}
	p.color = ColorSuccess
	p.Print()
}

func (p *Printable) PrintWarning() {
	if !DisableEmojis {
		p.text = "⚠️ " + p.text
	}
	p.color = ColorWarning
	p.Print()
}

func (p *Printable) PrintError() {
	if !DisableEmojis {
		p.text = "❕ " + p.text
	}
	p.color = ColorError
	p.Print()
	os.Exit(1)
}

func (p *Printable) PrintDebug() {
	if !DisableEmojis {
		p.text = "🐞 " + p.text
	}
	p.color = ColorDebug
	if PrintDebugMessages {
		p.Print()
	}
}
