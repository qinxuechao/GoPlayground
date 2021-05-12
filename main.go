package main

import (
	"GoPlayground/db"
	"GoPlayground/lib/log"
	"GoPlayground/model"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

func initArgs() *model.MysqlParams {
	var Params model.MysqlParams
	Params.Host = *flag.String("h", "127.0.0.1", "host ip, default is 127.0.0.1")
	Params.Port = *flag.String("p", "8080", "mysql port, default is 3306")
	Params.User = *flag.String("u", "root", "username, default is root")
	Params.Password = *flag.String("P", "password", "password, default is password")
	Params.Database = *flag.String("db", "test", "database, default is test database")

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

	rows, err := conn.Query("select * from user")
	if err != nil {
		logger.Errorf(err.Error())
	}

	var res model.User
	for rows.Next() {
		err = rows.Scan(&res.Id, &res.Name, &res.Department, &res.Role)
		fmt.Println(res)
	}

	querySql := "select * from user where id = 2"

	err = conn.QueryRow(querySql).Scan(&res.Id, &res.Name, &res.Department, &res.Role)
	if err != nil {
		// 如果是封装的方法直接返回
		// return errors.Wrapf(err,"execute querySql err: %s", querySql)

		// 主方法中直接打印忽略
		error := errors.Wrapf(err,"execute querySql err: %s", querySql)
		logger.Errorf("%+v\n", error)
	}
}
