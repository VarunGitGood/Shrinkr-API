package database

import (
	"api/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	rdb  *redis.Client
	rctx context.Context
)

func ConnectRedis() error {
	opt, err := redis.ParseURL(config.Config("REDIS_URL"))
	if err != nil {
		fmt.Println(err)
	}
	rdb = redis.NewClient(opt)
	rctx = context.Background()
	res := rdb.Ping(rctx)
	if res.Err() != nil {
		return res.Err()
	}
	fmt.Println("Connected to Redis")
	return nil
}

func StoreMapping(link *Link) error {
	result := rdb.HSet(rctx, "HASH", link.ShortURL, link.LongURL)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func GetMappings(username string) (map[string]string, error) {
	// result := rdb.HGetAll(rctx, username)
	// if result.Err() != nil {
	// 	return nil, result.Err()
	// }
	// return result.Val(), nil
	return nil, nil
}

func GetLongURL(shortURL string) (string, error) {
	result := rdb.HGet(rctx, "HASH", shortURL)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Val(), nil
}
