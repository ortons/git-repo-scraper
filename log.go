package main

import "fmt"

var Verbosity int

func logWarn(format string, a ...any) {
	fmt.Printf(format, a...)
	fmt.Println()
}

func logFmt(format string, a ...any) {
	if Verbosity > 1 {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
func logMsg(msg string) {
	if Verbosity > 1 {
		fmt.Println(msg)
	}
}

func logDebug(format string, a ...interface{}) {
	if Verbosity > 2 {
		{
			return
		}
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
