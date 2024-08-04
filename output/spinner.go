package cli

import (
	"fmt"
	"strings"
	"time"
)

type Spinner struct {
	_frame    int
	_message  string
	_stop     bool
	_time     int
	_showTime bool
}

func (s *Spinner) Run(message string, showTime bool) {
	s._stop = false
	s._frame = 0
	s._message = message
	s._time = 0
	s._showTime = showTime

	go s.draw()
}

func (s *Spinner) toString() string {
	frames := strings.Split("⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏", "")
	s._frame = (s._frame + 1) % len(frames)
	result := fmt.Sprintf("%s %s", frames[s._frame], s._message)
	if s._showTime {
		var formattedTime, _ = time.ParseDuration(fmt.Sprintf("%dms", s._time))
		return fmt.Sprintf("%s (%s elapsed)", result, formattedTime.Truncate(10000).String())
	}
	return result
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
