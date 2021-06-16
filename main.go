package main

import (
	"GoPlayground/lib/log"
	"context"

	"github.com/go-redis/redis/v8"
	logger "github.com/sirupsen/logrus"
)

var ctx = context.Background()
// ten bytes tests
const onebyte = "a"

func main() {
	// Init logs
	log.Init()
	// Init Redis
	logger.Info("init redis clinet...")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.Errorf("init redis client err: %+v", err)
		panic(err)
	}
	logger.Info("init client succeed")

	// remove all old data
	_, err := rdb.FlushAll(ctx).Result()
	if err != nil {
		logger.Errorf("flushall redis db err: %+v", err)
		panic(err)
	}
	logger.Info("flushed redis")

	info := rdb.Info(ctx)
	logger.Info(info)

	data := ""
	for i := 0; i < 50000; i++ {
		data += onebyte
	}

	for i := 0; i < 1000; i++ {
		err := rdb.Set(ctx, string(i), data, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	logger.Info("inserting completed")

	info = rdb.Info(ctx)
	logger.Info(info)
}
