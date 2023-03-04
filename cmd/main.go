package main

import (
	"github.com/devalexandre/mscli/internal/architecture"
	"github.com/devalexandre/mscli/internal/create"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := cobra.Command{Use: "ms"}
	rootCmd.AddCommand(architecture.Init())
	rootCmd.AddCommand(create.Init())

	rootCmd.Execute()
}
