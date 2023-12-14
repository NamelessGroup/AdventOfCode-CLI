package cli

import (
	"fmt"
	"strings"
	"time"
)

type Spinner struct {
	_frame   int
	_message string
	_stop    bool
}

func (s *Spinner) Run(message string) {
	s._stop = false
	s._frame = 0
	s._message = message

	go s.draw()
}

func (s *Spinner) toString() string {
	frames := strings.Split("⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏", "")
	s._frame = (s._frame + 1) % len(frames)
	return fmt.Sprintf("%s %s", frames[s._frame], s._message)
}

func (s *Spinner) draw() {
	for !s._stop {
		ToPrint(s.toString()).NewLine(false).Print()
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Spinner) Stop() {
	s._stop = true
	ToPrint("").NewLine(false).Print()
}

func (s *Spinner) Reprint() {
	ToPrint(s.toString()).NewLine(false).Print()
}
