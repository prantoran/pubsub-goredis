package psredis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Subscriber struct {
	client *redis.Client
	Pubsub *redis.PubSub
}

func (s *Subscriber) Connect(addr string) error {
	// e.g. addr = "localhost:6379"
	s.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Println("lis client:", s.client)
	_, err := s.client.Ping().Result()
	return err
}

func (s *Subscriber) Subscribe(chn ...string) error {
	s.Pubsub = s.client.Subscribe(chn...)
	// Wait for subscription to be created before publishing message.
	// remember to defer Pubsub.Close()
	subscr, err := s.Pubsub.ReceiveTimeout(time.Second)
	if err != nil {
		return err
	}
	log.Println("subscribe:", subscr)
	return nil
}

func (s *Subscriber) Listen(chn ...string) (*redis.Message, error) {
	msg, err := s.Pubsub.ReceiveMessage()
	return msg, err
}
