package languages

import "fmt"

type Test struct{}

func (t Test) GetSolveCommand(day int, task int) string {
	return "python3 ./test.py"
}

func (t Test) GetTestCommand(day int, task int) string {
	return fmt.Sprintf("echo This is the test command for day %d task %d", day, task)
}

func (t Test) GetPreparationCommand(day int, task int) string {
	return "echo This is never seen"
}
