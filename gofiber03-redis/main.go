package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gofiber03-redis/initialize"
)

func init() {
	//初始化redis连接
	initialize.Setup()
}

/*
利用SETNX实现分布式锁
SETNX 是『SET if Not eXists』(如果不存在，则 SET)的简写。
将 key 的值设为 value ，当且仅当 key 不存在，若给定的 key 已经存在，则 SETNX 不做任何动作。
如果setnx 返回ok 说明拿到了锁（这个变量过期了，当前线程设置了值）；
如果setnx 返回 nil，说明拿锁失败，被其他线程占用（这个key还存在）。

setnx key value:只有当key不存在的情况下，将key设置为value；若key存在，不做任何操作，结果成功返回1，失败返回0
setex key seconds value：将key值设置为value，并将设置key的生存周期,如果key值存在，使用setex将覆盖原有值

*/
func testNX() {

	// SET key value EX 10 NX
	//key已经存在返回false，key不存在设置成功返回true，设置过期时间
	//set, _ := initialize.RDB.SetNX(context.Background(), "key", "value", 10*time.Second).Result()

	// SET key value keepttl NX 设置过期时间 -1 永久不过期
	set, _ := initialize.RDB.SetNX(context.Background(), "key", "value", redis.KeepTTL).Result()
	fmt.Println(set)
}

func main() {
	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		testNX()

		////写值
		//err := initialize.RDB.Set(context.Background(), "password", "1234567", 0).Err()
		//if err != nil {
		//	panic(err)
		//}
		//
		////获取值
		//val, _ := initialize.RDB.Get(context.Background(), "name").Result()
		//if err == redis.Nil {
		//	fmt.Println("key does not exist")
		//} else if err != nil {
		//	panic(err)
		//}

		return ctx.SendString("hello world")
	})

	app.Listen(":8080")
}
