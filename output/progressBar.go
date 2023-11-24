package cli

import (
	"fmt"
)

type ProgressBar struct {
	_taskCount int
	_tasksDone int
	_message   string
}

func NewProgressBar(initialTaskCount int, initialMessage string) *ProgressBar {
	bar := &ProgressBar{
		_taskCount: initialTaskCount,
		_tasksDone: 0,
		_message:   initialMessage,
	}

	bar.draw()

	return bar
}

func (p *ProgressBar) getPercentage() float64 {
	return float64(p._tasksDone) / float64(p._taskCount)
}

func (p *ProgressBar) GotoNextTask(newTask string) {
	p._message = newTask
	if p._tasksDone < p._taskCount {
		p._tasksDone += 1
	}
}

func (p *ProgressBar) SetMessage(message string) {
	p._message = message

	p.draw()
}

func (p *ProgressBar) IncreaseDoneTasks() {
	if p._tasksDone < p._taskCount {
		p._tasksDone += 1
	}

	p.draw()
}

func (p *ProgressBar) SetDoneTasks(tasksDone int) {
	p._tasksDone = tasksDone

	p.draw()
}

func (p *ProgressBar) SetTotalTasks(totalTasks int) {
	p._taskCount = totalTasks

	p.draw()
}

func (p *ProgressBar) Finish(message string) {
	barString := p.getBarWithPercentage()

	ToPrint(barString).Color(ColorSuccess).NewLine(false).Print()
	ToPrint(message).Append(true).Print()
}

func (p *ProgressBar) Cancel(message string) {
	barString := p.getBarWithPercentage()

	ToPrint(barString).Color(ColorError).NewLine(false).Print()
	ToPrint(message).Append(true).Print()
}

func (p *ProgressBar) getBarWithPercentage() string {
	barWidth := 20
	bar := getFormattedBar(p.getPercentage(), barWidth)
	percentageFormated := fmt.Sprintf("%.2f%%", p.getPercentage()*100)

	spacesLength := 10 - len(percentageFormated)
	spaces := ""
	for i := 0; i < spacesLength; i++ {
		spaces += " "
	}

	return fmt.Sprintf("[%s] %s%s", bar, percentageFormated, spaces)
}

func (p *ProgressBar) draw() {
	barString := p.getBarWithPercentage()

	ToPrintf("%s%s", barString, p._message).NewLine(false).Print()
}

func getFormattedBar(percentage float64, barLength int) string {
	bars := []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}

	formattedBar := ""
	for i := 0; i < barLength; i++ {
		if int(percentage*float64(barLength)) > i {
			formattedBar += bars[len(bars)-1]
		} else if int(percentage*float64(barLength)) < i {
			formattedBar += bars[0]
		} else {
			formattedBar += bars[int((percentage*float64(barLength)-float64(i))*float64(len(bars)))]
		}
	}

	return formattedBar
}
