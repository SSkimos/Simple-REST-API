package utils

import (
	"github.com/nats-io/nats.go"
)

func InitNATSConnection(natsUrl string) (*nats.Conn, error) {
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func CloseConnection(nc *nats.Conn) {
	nc.Close()
}
