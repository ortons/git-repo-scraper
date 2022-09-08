package main

import (
	"fmt"
	"reflect"
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

func Test_buildShell(t *testing.T) {
	type args struct {
		dirs []gitDirEntry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildShell(tt.args.dirs)
		})
	}
}

func Test_createFallbackFilename(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFallbackFilename(tt.args.root); got != tt.want {
				t.Errorf("createFallbackFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createRepos(t *testing.T) {
	type args struct {
		repos []gitDirEntry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRepos(tt.args.repos)
		})
	}
}

func Test_dirPathWalk(t *testing.T) {
	type args struct {
		root   string
		filter string
	}
	tests := []struct {
		name    string
		args    args
		want    []gitDirEntry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dirPathWalk(tt.args.root, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("dirPathWalk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dirPathWalk() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_execCmd(t *testing.T) {
	type args struct {
		name string
		arg  []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execCmd(tt.args.name, tt.args.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("execCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("execCmd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exportCsv(t *testing.T) {
	type args struct {
		filename string
		repos    []gitDirEntry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exportCsv(tt.args.filename, tt.args.repos)
		})
	}
}

func Test_fileExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileExists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fileExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_log(t *testing.T) {
	type fields struct {
		Verbose []bool
		File    string
		Action  string
		Args    struct {
			RootFolder string
		}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := options{
				Verbose: tt.fields.Verbose,
				File:    tt.fields.File,
				Action:  tt.fields.Action,
				Args:    tt.fields.Args,
			}
			o.log()
		})
	}
}

func Test_parseGitDetails(t *testing.T) {
	type args struct {
		e *gitDirEntry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseGitDetails(tt.args.e)
		})
	}
}

func Test_readCsv(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []gitDirEntry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCsv(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCsv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkRoot(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want []gitDirEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := walkRoot(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walkRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
