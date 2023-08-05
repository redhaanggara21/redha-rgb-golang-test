package config

import "github.com/go-redis/redis"

var rdb *redis.Client

func RedisInit() error {
	db := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       2,  // database 2
	})

	rdb = db

	return nil
}

func RedisConnect() *redis.Client {
	return rdb
}
