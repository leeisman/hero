package mysql

import (
	"database/sql"
	"fmt"
	entsql "github.com/facebookincubator/ent/dialect/sql"
	"hero/configs"
	"hero/database/ent"
	"hero/pkg/logger"
	"time"
)

var (
	defaultDB *sql.DB
)

func Client() *ent.Client {
	dataSourceName := configs.Get("database.mysql_url")

	var (
		dbUser                 = "root"
		dbPwd                  = "m7dnbdA1y80ypOe4"
		instanceConnectionName = "money-money-279309:asia-northeast1:mdev"
		dbName                 = "hero_dev"
	)

	dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

	if configs.EnvPath != "local" {
		dataSourceName = dbURI
	}

	db, err := sql.Open("mysql", dataSourceName)
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
	db.SetConnMaxLifetime(time.Second)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)

	db2, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Error("failed connecting to mysql: " + err.Error())
		return nil
	}
	if db2 == nil {
		logger.Error("db nil")
		return nil
	}
	db2.SetMaxIdleConns(1)
	db2.SetMaxOpenConns(1)
	db2.SetConnMaxLifetime(time.Hour)
	defaultDB = db2
	return ent.NewClient(ent.Driver(drv))
}

func DB() *sql.DB {
	return defaultDB
}
