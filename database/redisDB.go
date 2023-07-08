package database

import (
	"api/config"
	"api/types"
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

func StoreMapping(link *types.LinkDTO) error {
	result := rdb.Set(rctx, link.ShortURL, link.LongURL, 0)
	if result.Err() != nil {
		fmt.Println(result.Err())
		return result.Err()
	}
	return nil
}

func GetLongURL(shortURL string) (string, error) {
	result := rdb.Get(rctx, shortURL)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Val(), nil
}
