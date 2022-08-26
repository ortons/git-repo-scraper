package main

import (
	"fmt"
	"testing"
)

func Test_parseFlags(t *testing.T) {
	args := []string{
		"-vv",
		"--offset=5",
		"-n", "Me",
		"-p", "3",
		"-s", "hello",
		"-s", "world",
		"--ptrslice", "hello",
		"--ptrslice", "world",
		"--intmap", "a:1",
		"--intmap", "b:5",
		"--filename", "hello.go",
		"id",
		"10",
		"remaining1",
		"remaining2",
	}

	fmt.Println(args)
}
