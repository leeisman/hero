package cmd

import (
	"github.com/spf13/cobra"
	"hero/routes"
)

var Server = &cobra.Command{
	Use:          "server",
	SilenceUsage: true,
	Short:        "Start API Server",
	Run: func(cmd *cobra.Command, args []string) {
		routes.Run()
	},
}
