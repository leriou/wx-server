package library

import (
	"github.com/go-redis/redis"
)

type Di struct{

}

var (
	redisClient *redis.Client
)

func (d *Di) GetRedis() *redis.Client{
	if redisClient != nil {
		return redisClient
	} 
	redisOptions := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	return redis.NewClient(redisOptions)
}