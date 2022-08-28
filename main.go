package main

import (
	"bytes"
	"fmt"
	"github.com/jessevdk/go-flags"
	fs "io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var Opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	File    string `short:"f" long:"file" description:"csv file for import / export operations. ignored for shell operations"`
	Action  string `short:"a" long:"action" description:"action [export | import | shell]" choice:"export" choice:"import" choice:"shell"`
	Args    struct {
		RootFolder string
	} `positional-args:"yes" required:"yes"`
}

func main() {

	Opts.Action = "read"

	_, err := flags.Parse(&Opts)

	if err != nil {
		panic(err)
	}
	Verbosity = len(Opts.Verbose)

	logInfo("File: %s\n", Opts.File)
	logInfo("Verbosity: %d\n", len(Opts.Verbose))
	logInfo("Action: %s\n", Opts.Action)
	logInfo("RootFolder: %s\n", Opts.Args.RootFolder)

	e, err := fileExists(Opts.Args.RootFolder)
	if !e || err != nil {
		logWarn("root directory %s is not valid: %s", Opts.Args.RootFolder, err)
		os.Exit(2)
	} else {
		if Verbosity > 1 {
			logInfo("root directory %s is valid", Opts.Args.RootFolder)
		}
	}
	dirs, err := dirPathWalk(Opts.Args.RootFolder, ".git")

	sort.Strings(dirs)
	switch Opts.Action {
	case "export":
		break
	case "import":
		break
	case "shell":
		break

	}
	buildShell(dirs)
}

func buildShell(dirs []string) {
	//git remote get-url origin
	n := len(dirs)
	for i := 0; i < n; i++ {
		dir := dirs[i]

		root := filepath.Dir(dir)

		cmd := exec.Command("git", "-C", dir, "remote", "get-url", "origin")
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err := cmd.Run()
		if err != nil {
			if Verbosity > 1 {
				fmt.Printf("%s: %s", dir, errb.String())
			}
		} else {
			parent := filepath.Dir(root)
			fmt.Printf("mkdir -p '%s' && cd '%s' && git clone %s", parent, parent, outb.String()) // strings.ReplaceAll(outb.String(), "\n", ""))
			if i < n-1 {
				fmt.Printf(" \\ ") //\r\n")
			} else {
				fmt.Println()
			}
		}
	}
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func dirPathWalk(root string, filter string) ([]string, error) {
	var folders []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && strings.EqualFold(d.Name(), filter) {
			folders = append(folders, path)
		}
		return nil
	})
	return folders, err
}
