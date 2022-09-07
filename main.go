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

type gitDirEntry struct {
	absDir    string `csv:"absDir"`
	relDir    string `csv:"relDir"`
	gitRemote string `csv:"gitRemote"`
	gitBranch string `csv:"gitBranch"`
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
		importCsv(opts.File)

		break
	case "shell":
		gitDirEntries := walkRoot(opts.Args.RootFolder)
		buildShell(gitDirEntries)
		break
	}

}

func (o options) log() {
	logInfo("file: %s\n", opts.File)
	logInfo("verbosity: %d\n", len(opts.Verbose))
	logInfo("action: %s\n", opts.Action)
	logInfo("rootFolder: %s\n", opts.Args.RootFolder)
}

func createFallbackFilename(root string) string {

	//just use the leaf-folder in the root, not full dir
	base := path.Base(root)
	return fmt.Sprintf("repos_%s_%s.csv", base, time.Now().Format("20060102_34"))

}

func exportCsv(filename string, dirs []gitDirEntry) {

	f, err := os.Create(filename)

	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range dirs {
		if err := w.Write(record.mkCsvEntry()); err != nil {
			logWarn("error writing record %s to file", err, record.gitBranch)
		}
	}

}

func (g *gitDirEntry) mkCsvEntry() []string {
	return []string{g.absDir, g.relDir, g.gitRemote, g.gitBranch}
}

func importCsv(file string) ([]gitDirEntry, error) {
	if len(file) == 0 {
		return nil, errors.New("file is required")
	}
	var results []gitDirEntry

	if e, err := fileExists(file); !e || err != nil {
		return nil, errors.New("file is not accessible")
	}

	if f, err := os.Open(file); err != nil {
		return nil, err
	} else {
		r := csv.NewReader(f)
		if records, err := r.ReadAll(); err != nil {
			log.Fatal(err)
		} else {

			for _, record := range records {
				results = append(results, gitDirEntry{record[0], record[1], record[2], record[3]})
			}
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

	logInfo("root directory %s is valid", root)

	dirs, err := dirPathWalk(root, ".git")
	// sort.Strings(dirs)
	return dirs
}

func dirPathWalk(root string, filter string) ([]gitDirEntry, error) {
	var gitDirEntries []gitDirEntry

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if d.IsDir() && strings.EqualFold(d.Name(), filter) {
			g := gitDirEntry{absDir: path,
				relDir: path,
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
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil || errb.Len() > 0 {
		return "", errors.New(errb.String())
	}

	return strings.TrimSuffix(outb.String(), "\n"), nil
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
