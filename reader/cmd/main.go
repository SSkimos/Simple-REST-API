package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"log"
	"reader/pkg/common/db"
	"reader/pkg/common/models"
	"reader/pkg/common/utils"
	"reader/pkg/employees"
	"time"
)

func main() {
	time.Sleep(8 * time.Second)
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)
	natsUrl := viper.Get("NATS_URL").(string)

	nc := utils.InitNatsConnection(natsUrl)
	database := db.InitDBConnection(dbUrl)
	defer utils.CloseConnection(nc)

	messageQueue := make(chan *nats.Msg)
	utils.CreateSubscription(nc, messageQueue)
	for {
		select {
		case messageFromChan := <-messageQueue:
			fmt.Printf("Received request from chan: %s\n", messageFromChan.Data)
			var employee models.Employee
			err := json.Unmarshal(messageFromChan.Data, &employee)
			if err != nil {
				log.Panic("Unmarshal error: %w")
			}
			//TODO: check?

			employees.AddEmployee(database, employee)
			utils.SendReply(nc, messageFromChan, "Message processed successfully")

			if err := nc.LastError(); err != nil {
				log.Panic(err)
			}
		}
	}
}
