package main

import "fmt"

type gitDirEntry struct {
	absDir    string `csv:"absDir"`
	relDir    string `csv:"relDir"`
	gitRemote string `csv:"gitRemote"`
	gitBranch string `csv:"gitBranch"`
}

func (g *gitDirEntry) mkCsvEntry() []string {
	return []string{g.absDir, g.relDir, g.gitRemote, g.gitBranch}
}

func (g *gitDirEntry) toString() string {
	v := fmt.Sprintf("%s, %s %s", g.absDir, g.gitRemote, g.gitBranch)
	return v
}
