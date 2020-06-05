package cmd

import (
	"github.com/spf13/cobra"
	"hero/migrations"
)

var Migrate = &cobra.Command{
	Use:          "migrate",
	SilenceUsage: true,
	Short:        "Database migrate",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Migrate()
	},
}
