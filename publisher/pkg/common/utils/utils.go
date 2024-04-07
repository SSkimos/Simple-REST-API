package utils

import (
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
