package main

import (
	"GoPlayground/lib/log"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
	)

var ctx = context.Background()


func main() {
	log.Init()
	logger.Infof("init redis clinet...")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
