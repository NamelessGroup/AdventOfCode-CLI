package utils

import "fmt"

type FileTemplate struct {
	Content  string
	Filename string
}

func GetChallengeDirectory(year int, day int) string {
	return fmt.Sprintf("./%d/%02d/", year, day)
}
