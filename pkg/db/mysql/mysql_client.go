package mysql

import (
	"database/sql"
	entsql "github.com/facebookincubator/ent/dialect/sql"
	"hero/configs"
	"hero/database/ent"
	"hero/pkg/logger"
	"time"
)

var (
	defaultClient *ent.Client
)

func Client() *ent.Client {
	localDataSourceName := configs.Get("database.mysql_url")
	db, err := sql.Open("mysql", localDataSourceName)
	if err != nil {
		logger.Error("failed connecting to mysql: " + err.Error())
		return nil
	}
	if db == nil {
		logger.Error("db nil")
		return nil
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv))
}
