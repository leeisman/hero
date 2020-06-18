package mysql

import (
	"database/sql"
	"fmt"
	entsql "github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"hero/configs"
	"hero/database/ent"
	"hero/pkg/logger"
	"os"
	"strconv"
	"time"
)

var (
	defaultDB     *sql.DB
	defaultClient *ent.Client
)

func init() {
	dataSourceName := configs.Get("database.mysql_url")

	var (
		dbUser                 = "root"
		dbPwd                  = "m7dnbdA1y80ypOe4"
		instanceConnectionName = "34.84.24.127"
		dbName                 = "hero_dev"
	)

	//dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPwd, instanceConnectionName, 3306, dbName)
	dbMaxIdleConns := 1
	maxOpenConns := 5
	if true {
		dataSourceName = dbURI
		if MAX_IDLE_CONNS := os.Getenv("MAX_IDLE_CONNS"); MAX_IDLE_CONNS != "" {
			dbMaxIdleConns, _ = strconv.Atoi(MAX_IDLE_CONNS)
		}
		if MAX_OPEN_CONNS := os.Getenv("MAX_OPEN_CONNS"); MAX_OPEN_CONNS != "" {
			maxOpenConns, _ = strconv.Atoi(MAX_OPEN_CONNS)
		}
	}

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic("failed connecting to mysql: " + err.Error())
	}
	if db == nil {
		panic("db nil")
	}

	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(1 * time.Second)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)

	defaultDB = db
	defaultClient = ent.NewClient(ent.Driver(drv)).Debug()
}

func Client() *ent.Client {
	return defaultClient
}

func DB() *sql.DB {
	return defaultDB
}

func Rollback(tx *ent.Tx, err error) {
	logger.Print("mysql rollback before err", err.Error())
	if err := tx.Rollback(); err != nil {
		logger.Print("mysql rollback err", err.Error())
	}
}

func Close() error {
	return defaultDB.Close()
}
