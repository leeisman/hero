package cmd

import (
	"github.com/facebookincubator/ent/examples/start/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"hero/logger"
)

var Migrate = &cobra.Command{
	Use:          "migrate",
	SilenceUsage: true,
	Short:        "Database migrate",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ent.Open("mysql", "root:root@127.0.0.1:3306)/hero?parseTime=True")
		if err != nil {
			logger.Error(err.Error())
		}
		defer client.Close()
	},
}
