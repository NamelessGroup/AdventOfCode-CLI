package cli

import (
	"fmt"
	"os"
)

func Error(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func Warn(message string) {
	fmt.Println(message)
}

func Log(message string) {
	fmt.Println(message)

}

func Debug(message string) {
	fmt.Println(message)
}

func Success(message string) {
	fmt.Println(message)
}