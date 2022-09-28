package utils

import "fmt"

var Verbosity int

func LogWarn(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println()
}

func LogFmt(format string, a ...interface{}) {
	if Verbosity > 1 {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
func LogMsg(msg string) {
	if Verbosity > 1 {
		fmt.Println(msg)
	}
}

func LogDebug(format string, a ...interface{}) {
	if Verbosity > 2 {
		{
			return
		}
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
