package main

import "fmt"

var Verbosity int

func logWarn(format string, a ...interface{}) {
	fmt.Printf(format, a...)

}

func logInfo(format string, a ...interface{}) {
	if Verbosity < 0 {
		return
	}
	fmt.Printf(format, a...)

}

func logDebug(format string, a ...interface{}) {
	if Verbosity < 1 {
		return
	}
	fmt.Printf(format, a...)

}
