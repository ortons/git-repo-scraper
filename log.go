package main

import "fmt"

var Verbosity int

func logWarn(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println()
}

func logInfo(format string, a ...interface{}) {
	if Verbosity > 1 {
		fmt.Printf(format, a...)
		fmt.Println()
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
