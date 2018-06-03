package psredis

import (
	"github.com/go-redis/redis"
	"magic.pathao.com/fellowship/rides-simple-dispatch/conf"
)

var rClient *redis.Client

func ConnectToRedis() error {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: conf.GetApp().Redis.Password, // no password set
		DB:       conf.GetApp().Redis.DB,       // use default DB
	})
	_, err := rClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func DefaultRedis() *redis.Client {
	return rClient
}
