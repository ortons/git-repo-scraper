package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	fs "io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	File    string `short:"f" long:"file" description:"csv file for import / export operations. ignored for shell operations"`
	Action  string `short:"a" long:"action" description:"action [export | import | shell]" choice:"export" choice:"import" choice:"shell"`
	Args    struct {
		RootFolder string
	} `positional-args:"yes" required:"yes"`
}

var opts options

func main() {

	opts.Action = "shell"

	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
	}
	Verbosity = len(opts.Verbose)
	opts.log()

	switch opts.Action {
	case "export":
		var file = opts.File
		if len(file) == 0 {
			file = createFallbackFilename(opts.Args.RootFolder)
		}
		gitDirEntries := walkRoot(opts.Args.RootFolder)
		exportCsv(file, gitDirEntries)
		break
	case "import":
		if r, err := readCsv(opts.File); err != nil {
			log.Fatal(err)
		} else {
			createRepos(r)
		}
		break
	case "shell":
		gitDirEntries := walkRoot(opts.Args.RootFolder)
		buildShell(gitDirEntries)
		break
	}

}

func createRepos(repos []gitDirEntry) {
	for _, repo := range repos {
		logMsg(repo.toString())

		if err := os.Mkdir(repo.absDir, os.ModePerm); err != nil {
			logWarn("unable to create the target directory: %s", err.Error())
		}

		var out, err = execCmd("git", "clone", "--branch", repo.gitBranch, repo.gitRemote, repo.absDir)
		if err != nil {
			logWarn("failed to clone git repo %s:%s to %s. %s", repo.gitRemote, repo.gitBranch, repo.absDir, err)
		}
		logMsg(out)
	}
}

func (o options) log() {
	logFmt("file: %s\n", opts.File)
	logFmt("verbosity: %d\n", len(opts.Verbose))
	logFmt("action: %s\n", opts.Action)
	logFmt("rootFolder: %s\n", opts.Args.RootFolder)
}

func createFallbackFilename(root string) string {

	//just use the leaf-folder in the root, not full dir
	base := path.Base(root)
	return fmt.Sprintf("repos_%s_%s.csv", base, time.Now().Format("20060102_34"))

}

func exportCsv(filename string, repos []gitDirEntry) {

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range repos {
		if err := w.Write(record.mkCsvEntry()); err != nil {
			logWarn("error writing record %s to file", err, record.gitBranch)
		}
	}
}

func readCsv(file string) ([]gitDirEntry, error) {
	if len(file) == 0 {
		return nil, errors.New("file is required")
	}
	var results []gitDirEntry

	if e, err := fileExists(file); !e || err != nil {
		return nil, errors.New("file is not accessible")
	}

	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	if records, err := r.ReadAll(); err != nil {
		log.Fatal(err)
	} else {
		for _, record := range records {
			results = append(results, gitDirEntry{record[0], record[1], record[2], record[3]})
		}
	}

	return results, nil
}

func buildShell(dirs []gitDirEntry) {
	//git remote get-url origin
	for _, dir := range dirs {
		fmt.Printf("mkdir -p '%s' && cd '%s' && git clone %s -b %s", dir.absDir, dir.absDir, dir.gitRemote, dir.gitBranch)
	}
}

func walkRoot(root string) []gitDirEntry {
	exists, err := fileExists(opts.Args.RootFolder)
	if !exists || err != nil {
		log.Fatalf("root directory %s is not valid: %s\n", root, err)
	}

	logFmt("root directory %s is valid", root)

	dirs, err := dirPathWalk(root, ".git")
	// sort.Strings(dirs)
	return dirs
}

func dirPathWalk(root string, filter string) ([]gitDirEntry, error) {
	var gitDirEntries []gitDirEntry

	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {

		if d.IsDir() && strings.EqualFold(d.Name(), filter) {
			parent := path.Dir(p)
			g := gitDirEntry{absDir: parent,
				relDir: parent,
			}
			parseGitDetails(&g)

			gitDirEntries = append(gitDirEntries, g)
		}
		return nil
	})
	return gitDirEntries, err
}

func parseGitDetails(e *gitDirEntry) {

	//get remote
	var out, err = execCmd("git", "-C", e.absDir, "remote", "get-url", "origin")
	if err != nil {
		logWarn("failed to get git remote %s: %s", e.absDir, err)
	}
	(*e).gitRemote = out

	//get branch
	out, err = execCmd("git", "-C", e.absDir, "branch", "--show-current")

	if err != nil {
		logWarn("failed to get git branch %s: %s", e.absDir, err)
	}

	(*e).gitBranch = out

}

func execCmd(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out, errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil || errs.Len() > 0 {
		return "", errors.New(errs.String())
	}

	return strings.TrimSuffix(out.String(), "\n"), nil
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
