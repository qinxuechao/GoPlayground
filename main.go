package main

import (
	"GoPlayground/db"
	"GoPlayground/lib/log"
	"GoPlayground/model"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
)

func initArgs() *model.MysqlParams {
	var Params model.MysqlParams
	Params.Host = *flag.String("h", "127.0.0.1", "host ip, default is 127.0.0.1")
	Params.Port = *flag.String("p", "3306", "mysql port, default is 3306")
	Params.User = *flag.String("u", "root", "username, default is root")
	Params.Password = *flag.String("P", "password", "password, default is password")
	Params.Password = *flag.String("db", "test", "database, default is test database")

	flag.Parse()

	return &Params
}

func main() {
	log.Init()
	logger.Infof("starting service...")
	logger.Infof("parsing args...")
	params := initArgs()
	conn := db.Init(params)


	defer conn.Close()

	insert, err := conn.Query("INSERT INTO test.user VALUES ( 1, 'xiaoming', 'r&d', 'engineer')")
	if err != nil {
		logger.Errorf(err.Error())
	}

	defer insert.Close()
}
