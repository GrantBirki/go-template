package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	var name string

	rootCmd := &cobra.Command{
		Use:   "go-template",
		Short: "A simple CLI template built with Cobra",
		Long: `A simple CLI template built with Cobra.

This is a template project for building CLI applications in Go using the Cobra library.
You can use this as a starting point for your own CLI applications.`,
		Run: func(cmd *cobra.Command, args []string) {
			if name != "" {
				cmd.Printf("Hello %s!\n", name)
			} else {
				cmd.Printf("Hello World!\n")
			}
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet (optional)")
	return rootCmd
}

// Run executes the CLI with explicit inputs and outputs.
func Run(args []string, stdout, stderr io.Writer) error {
	rootCmd := newRootCmd()
	rootCmd.SetArgs(args)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	return rootCmd.Execute()
}
