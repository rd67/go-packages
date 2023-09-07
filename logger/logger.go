package logger

import "fmt"

var Version string = "1.0"

func LogInfo(message string) {
	fmt.Println("[LogInfo]: "+ message)
}
