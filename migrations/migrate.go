package migrations

import (
	"github.com/facebookincubator/ent/examples/start/ent"
	_ "github.com/go-sql-driver/mysql"
	"hero/logger"
)

func Migrate() {
	client, err := ent.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hero?parseTime=True")
	if err != nil {
		logger.Error(err.Error())
	}
	defer client.Close()
}
