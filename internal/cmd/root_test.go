package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "default hello world",
			args:     []string{},
			expected: "Hello World!\n",
		},
		{
			name:     "hello with name flag",
			args:     []string{"--name", "John"},
			expected: "Hello John!\n",
		},
		{
			name:     "hello with short name flag",
			args:     []string{"-n", "Alice"},
			expected: "Hello Alice!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)
			err := Run(tt.args, stdout, stderr)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if stdout.String() != tt.expected {
				t.Fatalf("expected output %q, got %q", tt.expected, stdout.String())
			}
			if stderr.Len() != 0 {
				t.Fatalf("expected no stderr, got %q", stderr.String())
			}
		})
	}
}

func TestRunHelp(t *testing.T) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	err := Run([]string{"--help"}, stdout, stderr)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr, got %q", stderr.String())
	}
	output := stdout.String()

	mustContain := []string{
		"A simple CLI template built with Cobra",
		"Usage:",
		"go-template [flags]",
		"Flags:",
		"-h, --help",
		"-n, --name string",
	}

	for _, substr := range mustContain {
		if !strings.Contains(output, substr) {
			t.Fatalf("expected help output to contain %q", substr)
		}
	}
}

func TestRunInvalidFlag(t *testing.T) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	err := Run([]string{"--not-a-flag"}, stdout, stderr)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "unknown flag") {
		t.Fatalf("expected unknown flag error, got %q", err)
	}
}
