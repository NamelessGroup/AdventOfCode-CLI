package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
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

func (s *Spinner) draw() {
	frames := strings.Split("⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏", "")

	for !s._stop {
		s._frame = (s._frame + 1) % len(frames)
		Print(fmt.Sprintf("%s %s", frames[s._frame], s._message), color.BgBlack, Format{}, false)
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Spinner) Stop() {
	s._stop = true
	Print("", color.BgBlack, Format{}, false)
}
