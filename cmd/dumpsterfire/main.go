package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/DumpsterFIRE/dumpsterfire/internal/ui"
)

func main() {
	root := &cobra.Command{
		Use:   "dumpsterfire",
		Short: "DumpsterFIRE incident response workflows",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if ctx == nil {
				ctx = context.Background()
			}

			program := tea.NewProgram(ui.NewModel(ctx), tea.WithAltScreen())
			if err := program.Start(); err != nil {
				return err
			}
			return nil
		},
	}

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
