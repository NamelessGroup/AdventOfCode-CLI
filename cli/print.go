package cli

import (
	"fmt"
	"os"
)

func PrintError(message string) {
	fmt.Println(message)
	os.Exit(1)
}


func PrintWarning(message string) {
	fmt.Println(message)
}

func PrintLog(message string) {
	fmt.Println(message)

}

func PrintDebug(message string) {
	fmt.Println(message)
}

func PrintSuccess(message string) {
	fmt.Println(message)
}