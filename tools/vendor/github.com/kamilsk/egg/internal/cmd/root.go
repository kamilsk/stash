package cmd

import (
	"github.com/kamilsk/egg/internal/cmd/deps"
	"github.com/kamilsk/egg/internal/cmd/make"
	"github.com/kamilsk/egg/internal/cmd/tools"
	"github.com/kamilsk/egg/internal/cmd/vanity"
	"github.com/spf13/cobra"
)

// New returns the new root command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "egg",
		Short: "Extended go get",
		Long:  "Extended go get - alternative for the standard `go get` with a few little but useful features.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	cmd.AddCommand(deps.New(), make.New(), tools.New(), vanity.New())
	return cmd
}
