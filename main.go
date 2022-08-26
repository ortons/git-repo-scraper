package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
)

var Opts struct {
	Verbose   []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Operation string `short:"o" long:"operation" description:"operation [read | create]" choice:"read" choice:"create"`
	Args      struct {
		RootFolder string
	} `positional-args:"yes" required:"yes"`
}

func main() {

	Opts.Operation = "read"

	_, err := flags.Parse(&Opts)

	if err != nil {
		panic(err)
	}
	Verbosity := len(Opts.Verbose)
	fmt.Printf("Verbosity: %d\n", len(Opts.Verbose))
	fmt.Printf("Operation: %s\n", Opts.Operation)
	fmt.Printf("RootFolder: %s\n", Opts.Args.RootFolder)

	// cmd := exec.Command("tr", "a-z", "A-Z")
	//	out, err := cmd := exec.Command("find", "Opts.Args.RootFolder", "-path '*/.git/config'", "-execdir git remote get-url origin \\;").Output()

	e, err := exists(Opts.Args.RootFolder)
	if !e || err != nil {
		fmt.Sprintf("root directory %s is not valid/r/n", Opts.Args.RootFolder)
	}
	files, err := FilePathWalk(Opts.Args.RootFolder)

	if Verbosity > 0 {
		for _, file := range files {
			fmt.Println(file)
		}
	}
}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func FilePathWalk(root string) ([]string, error) {
	var folders []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			folders = append(folders, path)
		}
		return nil
	})
	return folders, err
}
