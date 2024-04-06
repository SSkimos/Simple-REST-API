package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"log"
	"reader/pkg/common/utils"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	natsUrl := viper.Get("NATS_URL").(string)
	messageQueue := make(chan *nats.Msg)

	nc, err := utils.InitConnection(natsUrl)
	if err != nil {
		log.Fatal("Connection to NATS error: %w", err)
		return
	}
	defer utils.CloseConnection(nc)

	utils.CreateSubscription(nc, messageQueue)
	for {
		select {
		case messageFromChan := <-messageQueue:
			fmt.Printf("Received request from chan: %s\n", messageFromChan.Data)
			utils.SendReply(nc, messageFromChan, "Message processed successfully")
			if err := nc.LastError(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
