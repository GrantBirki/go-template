package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			// Reset the name variable for each test
			name = ""

			// Capture output
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)

			// Set args and execute
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestRootCmdHelp(t *testing.T) {
	// Reset the name variable
	name = ""

	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Test help command
	rootCmd.SetArgs([]string{"--help"})
	err := rootCmd.Execute()

	assert.NoError(t, err)
	output := buf.String()

	// Check that help output contains expected elements
	assert.Contains(t, output, "A simple CLI template built with Cobra")
	assert.Contains(t, output, "Usage:")
	assert.Contains(t, output, "go-template [flags]")
	assert.Contains(t, output, "Flags:")
	assert.Contains(t, output, "-h, --help")
	assert.Contains(t, output, "-n, --name string")
}
