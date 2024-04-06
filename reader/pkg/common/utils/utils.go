package utils

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func InitNatsConnection(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal("Connection to NATS error: %w", err)
	}
	return nc
}

func CloseConnection(nc *nats.Conn) {
	nc.Close()
}

func CreateSubscription(nc *nats.Conn, messageQueue chan *nats.Msg) {
	fmt.Println("Subscriber is running, waiting for messages...")
	nc.Subscribe("new_employee", func(msg *nats.Msg) {
		// Получение и обработка сообщения от Publisher
		messageQueue <- msg
	})
}

func SendReply(nc *nats.Conn, msg *nats.Msg, reply string) {
	nc.Publish(msg.Reply, []byte(reply))
	nc.Flush()
}
