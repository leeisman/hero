package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	commander = &cobra.Command{
		SilenceUsage: true,
	}
)

func Execute() {
	commander.AddCommand(Server)
	if err := commander.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
