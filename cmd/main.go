package main

import (
	"fmt"
	"log"

	"github.com/prantoran/pubsub-goredis/pubsub/psredis"
)

func main() {
	lis := psredis.Subscriber{}
	if err := lis.Connect("localhost:6379"); err != nil {
		log.Println("lis con:", err)
		return
	}
	if err := lis.Subscribe("chan1", "chan2"); err != nil {
		log.Println("lis con:", err)
		return
	}
	pub := psredis.Publisher{}
	if err := pub.Connect("localhost:6379"); err != nil {
		log.Println("lis con:", err)
		return
	}

	go func() {
		for {
			select {
			case cmsg := <-lis.Pubsub.Channel():
				fmt.Println("csmg channnel:", cmsg.Channel, " pattern:", cmsg.Pattern,
					" payload:", cmsg.Payload, "\ncsmg.String():", cmsg.String())
			}
		}
	}()

	if err := pub.Publish("chan1", "first msg"); err != nil {
		log.Println("pub:", err)
		return
	}

	if err := pub.Publish("chan2", "second msg"); err != nil {
		log.Println("pub:", err)
		return
	}

	lis.Pubsub.Unsubscribe("chan1")

	forever := make(chan bool)

	<-forever
}
