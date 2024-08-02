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
	_time    int
}

func (s *Spinner) Run(message string) {
	s._stop = false
	s._frame = 0
	s._message = message
	s._time = 0

	go s.draw()
}

func (s *Spinner) toString() string {
	frames := strings.Split("⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏", "")
	s._frame = (s._frame + 1) % len(frames)
	var formattedTime, _ = time.ParseDuration(fmt.Sprintf("%dms", s._time))
	return fmt.Sprintf("%s %s (%s elapsed)", frames[s._frame], s._message, formattedTime.Truncate(10000).String())
}

func (s *Spinner) draw() {
	for !s._stop {
		ToPrint(s.toString()).NewLine(false).Print()
		s._time += 100
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
