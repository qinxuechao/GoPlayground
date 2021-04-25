package db

import (
	"GoPlayground/model"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Init(params *model.MysqlParams) *sql.DB {
	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", params.User, params.Password, params.Host,
		params.Port, params.Database)
	log.Infof("connecting to mysql db: %s:%s", params.Host, params.Port)
	db, err := sql.Open("mysql", dbPath)

	if err != nil {
		log.Fatal("connecting to db error: ",err.Error()) //退出程序
	}

	log.Infof("succeeded")
	return db
}
