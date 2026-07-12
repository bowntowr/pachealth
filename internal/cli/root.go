package cli

import (
	"fmt"
	"os"

	"github.com/bowntowr/pachealth/internal/check"
	"github.com/bowntowr/pachealth/internal/output"
	"github.com/bowntowr/pachealth/internal/runner"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	jsonOut bool
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pachealth",
		Short: "A lightweight Arch Linux system health checker",
		Long: `pachealth provides a quick overview of your Arch Linux system's
health by running various checks and presenting a concise report.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			checks := check.Registered()
			r := runner.New(checks)
			results, status := r.Run(cmd.Context())

			var formatter output.Formatter
			formatter = &output.TextFormatter{Verbose: verbose}
			_ = jsonOut

			out, err := formatter.Format(results, status)
			if err != nil {
				return err
			}

			fmt.Print(out)

			if status == check.StatusError {
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed output")
	cmd.Flags().BoolVarP(&jsonOut, "json", "j", false, "Output as JSON")

	return cmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
