package config

import (
	"backend/global"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func initRedis() {
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	// 新建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})

	// 检查Redis服务是否可用
	ctx := context.Background()
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Println("Failed to connect redis: ", err)
		global.RedisDB = nil
		return
	}

	global.RedisDB = redisClient
	log.Println("Redis initialized successfully")
}
