package cli

import (
	"fmt"

	"github.com/fatih/color"
)

type ProgressBar struct {
	_percentage float64
	_message    string
}

func (p *ProgressBar) Run(message string) {
	p._message = message
	p._percentage = 0

	go p.draw(color.FgWhite, false)
}

func (p *ProgressBar) Set(message string, percentage float64) {
	p._message = message
	p.SetPercentage(percentage)
}

func (p *ProgressBar) SetPercentage(percentage float64) {
	p._percentage = min(1, max(0, percentage))

	go p.draw(color.FgWhite, false)
}

func (p *ProgressBar) SetMessage(message string) {
	p._message = message

	go p.draw(color.FgWhite, false)
}

func (p *ProgressBar) Finish(message string) {
	p._percentage = 1
	p._message = message

	p.draw(color.FgGreen, true)
}

func (p *ProgressBar) Cancel(message string) {
	p.draw(color.FgRed, true)
}

func (p *ProgressBar) draw(col color.Attribute, breakAtEnd bool) {
	barWidth := 20
	bar := ""
	for i := 0; i < barWidth; i++ {
		bar += getTenthPercentage(p._percentage, i, barWidth)
	}

	percentageFormated := fmt.Sprintf("%.2f%%", p._percentage*100)
	spacesLength := 10 - len(percentageFormated)
	spaces := ""
	for i := 0; i < spacesLength; i++ {
		spaces += " "
	}

	Print(fmt.Sprintf("[%s] %s%s%s", bar, percentageFormated, spaces, p._message), col, Format{}, breakAtEnd)
}

func getTenthPercentage(percentage float64, index int, barLength int) string {
	bars := []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}

	// [12345678]
	if int(percentage*float64(barLength)) > index {
		return bars[len(bars)-1]
	} else if int(percentage*float64(barLength)) < index {
		return bars[0]
	} else {
		return bars[int((percentage*float64(barLength)-float64(index))*float64(len(bars)))]
	}
}
