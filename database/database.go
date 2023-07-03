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

func ConnectMongo() error {
	return nil
}

func CreateUser(user *User) error {
	result := rdb.HSet(rctx, "users", user.Username, user.Password)
	if result.Err() != nil {
		return result.Err()
	}
	fmt.Println(result.Val())
	return nil
}

func GetAllUsers() (map[string]string, error) {
	result := rdb.HGetAll(rctx, "users")
	if result.Err() != nil {
		return nil, result.Err()
	}
	fmt.Println(result.Val())
	return result.Val(), nil
}

func StoreMapping(link *Link, username string) error {
	// storing the link in the database as a hash with the username as the name of the hashmap and the short url as the key
	result := rdb.HSet(rctx, username, link.ShortURL, link.LongURL)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func GetMappings(username string) (map[string]string, error) {
	result := rdb.HGetAll(rctx, username)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}
