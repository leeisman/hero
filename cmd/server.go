package cmd

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"hero/database/ent/migrate"
	"hero/pkg/db/mysql"
	"hero/pkg/graceful"
	"hero/pkg/logger"
	"hero/routes"
	"time"
)

var Server = &cobra.Command{
	Use:          "server",
	SilenceUsage: true,
	Short:        "Start API Server",
	Run: func(cmd *cobra.Command, args []string) {
		//建立table
		AutoMigrate()
		graceful.Launch(routes.Run, routes.Shutdown, 1*time.Minute)
	},
}

func AutoMigrate() {
	client := mysql.Client()
	//defer client.Close()
	ctx := context.Background()
	// Run migration.
	err := client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.Error("failed creating schema resources: " + err.Error())
	}
	logger.Debug("auto migrate successful")
}
