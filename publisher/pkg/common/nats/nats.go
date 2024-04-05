package nats

import (
	"github.com/nats-io/nats.go"
	"log"
)

func InitConnection(url string) *nats.Conn {
	opts := []nats.Option{nats.Name("NATS Publisher")}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Panicf("Connection error: %v", err)
	}
	defer nc.Close()

	return nc
}
