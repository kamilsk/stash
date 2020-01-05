package main

import (
	"os"

	"github.com/kamilsk/egg/internal/cmd"
	"go.octolab.org/toolkit/cli/cobra"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

func main() {
	root := cmd.New()
	root.SetOut(os.Stdout)
	root.SetErr(os.Stderr)
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(version, date, commit))
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
