package cmd

import (
	"github.com/spf13/cobra"
	"hero/database/seeds"
)

var Seeder = &cobra.Command{
	Use:          "seed run",
	SilenceUsage: true,
	Short:        "run seed",
	Run: func(cmd *cobra.Command, args []string) {
		seeds.ActiveSeed()
	},
}
