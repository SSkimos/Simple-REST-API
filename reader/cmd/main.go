package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	time.Sleep(8 * time.Second)
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	//port := viper.Get("PORT").(string)
	//dbUrl := viper.Get("DB_URL").(string)
	//
	//h := db.Init(dbUrl)

	// Подключение к серверу NATS
	nc, err := nats.Connect("stan-nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Обработка сообщений из "my-subject"
	nc.Subscribe("new_employee", func(msg *nats.Msg) {
		// Получение и обработка сообщения от Publisher
		request := string(msg.Data)
		fmt.Printf("Received request: %s\n", request)

		// Отправка ответа обратно Publisher
		response := []byte("Message processed successfully")
		nc.Publish(msg.Reply, response)
	})

	// Ожидание сообщений
	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Subscriber is running, waiting for messages...")
	select {} // Бесконечное ожидание сообщений
}
