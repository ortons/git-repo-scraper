package utils

import "testing"

func Test_logDebug(t *testing.T) {
	type args struct {
		format string
		a      []any
	}
	Verbosity = 3

	var tests = []struct {
		name string
		args args
	}{
		{"Test", args{"debug", []any{"123", 123, "yury"}}},
		{"Test2", args{"debug %s %s %s", []any{"123", 123, "yury"}}},
		{"Test2", args{"debug %d %d %s", []any{"123", 123, "yury"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogDebug(tt.args.format, tt.args.a...)
		})
	}
}

func Test_logInfo(t *testing.T) {
	type args struct {
		format string
		a      []any
	}
	Verbosity = 2

	tests := []struct {
		name string
		args args
	}{{"Test", args{"info", []any{"123", 123, "yury"}}},
		{"Test2", args{"info %s %s %s", []any{"123", 123, "yury"}}},
		{"Test2", args{"info %d %d %s", []any{"123", 123, "yury"}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogFmt(tt.args.format, tt.args.a...)
		})
	}
}

func Test_logWarn(t *testing.T) {
	type args struct {
		format string
		a      []any
	}
	Verbosity = 2
	tests := []struct {
		name string
		args args
	}{{"Test", args{"warn", []any{"123", 123, "yury"}}},
		{"Test2", args{"warn %s %s %s", []any{"123", 123, "yury"}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogWarn(tt.args.format, tt.args.a...)
		})
	}
}
