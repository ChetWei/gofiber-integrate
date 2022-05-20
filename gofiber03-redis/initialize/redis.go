package initialize

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func Setup() {

	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

}
