package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// createTestRootCmd creates a fresh root command for testing to avoid state interference
func createTestRootCmd() *cobra.Command {
	var testName string

	cmd := &cobra.Command{
		Use:   "go-template",
		Short: "A simple CLI template built with Cobra",
		Long: `A simple CLI template built with Cobra.

This is a template project for building CLI applications in Go using the Cobra library.
You can use this as a starting point for your own CLI applications.`,
		Run: func(cmd *cobra.Command, args []string) {
			if testName != "" {
				cmd.Printf("Hello %s!\n", testName)
			} else {
				cmd.Printf("Hello World!\n")
			}
		},
	}

	cmd.Flags().StringVarP(&testName, "name", "n", "", "Name to greet (optional)")
	return cmd
}

func TestRootCmd(t *testing.T) {
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
			// Create a fresh command for each test
			cmd := createTestRootCmd()

			// Capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			// Set args and execute
			cmd.SetArgs(tt.args)
			err := cmd.Execute()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if buf.String() != tt.expected {
				t.Fatalf("expected output %q, got %q", tt.expected, buf.String())
			}
		})
	}
}

func TestRootCmdHelp(t *testing.T) {
	// Create a fresh command for the help test
	cmd := createTestRootCmd()

	// Capture output
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	// Test help command
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	output := buf.String()

	// Check that help output contains expected elements
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
