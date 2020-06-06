package cmd

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"hero/database/ent/migrate"
	"hero/pkg/db"
	"hero/pkg/logger"
)

func AutoMigrate() {
	client := db.Client()
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.Error("failed creating schema resources: " + err.Error())
	}
	logger.Debug("auto migrate successful")
}
