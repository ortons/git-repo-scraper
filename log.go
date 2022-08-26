package main

import "fmt"

var Verbosity int

func logWarn(format string, a ...any) {
	fmt.Printf(format, a...)

}

func logInfo(format string, a ...any) {
	if Verbosity < 0 {
		return
	}
	fmt.Printf(format, a...)

}

func logDebug(format string, a ...any) {
	if Verbosity < 1 {
		return
	}
	fmt.Printf(format, a...)

}
