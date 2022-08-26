package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	fs "io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
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
	Verbosity = len(Opts.Verbose)
	logInfo("Verbosity: %d\n", len(Opts.Verbose))
	logInfo("Operation: %s\n", Opts.Operation)
	logInfo("RootFolder: %s\n", Opts.Args.RootFolder)

	// cmd := exec.Command("tr", "a-z", "A-Z")
	//	out, err := cmd := exec.Command("find", "Opts.Args.RootFolder", "-path '*/.git/config'", "-execdir git remote get-url origin \\;").Output()

	e, err := fileExists(Opts.Args.RootFolder)
	if !e || err != nil {
		logWarn("root directory %s is not valid: %s", Opts.Args.RootFolder, err)
	} else {
		logInfo("root directory %s is valid", Opts.Args.RootFolder)
	}
	dirs, err := dirPathWalk(Opts.Args.RootFolder, ".git")

	sort.Strings(dirs)
	printRemotes(dirs)
}

func printRemotes(dirs []string) {
	//git remote get-url origin

	for _, dir := range dirs {

		out, err := exec.Command("git", "-C", dir, "remote", "get-url", "origin").Output()
		if err != nil {
			log.Fatal(err)
		}
		/*fmt.Printf("folder: %s, remote: %s", dir, out)
		 */
		fmt.Printf("git clone %s", out)

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
