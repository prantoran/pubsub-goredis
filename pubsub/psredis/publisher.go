package psredis

import (
	"github.com/go-redis/redis"
)

type Publisher struct {
	client *redis.Client
}

func (p *Publisher) Connect(addr string) error {
	// e.g. addr = "localhost:6379"
	p.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := p.client.Ping().Result()
	return err
}

func (p *Publisher) Publish(chn, msg string) error {
	return p.client.Publish(chn, msg).Err()
}
