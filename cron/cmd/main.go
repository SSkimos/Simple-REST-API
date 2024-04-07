package main

import (
	"cron/pkg/common/db"
	"fmt"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	time.Sleep(8 * time.Second)
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)

	fmt.Println("get config: %w", dbUrl)

	database := db.InitDBConnection(dbUrl)

	c := cron.New()
	c.AddFunc("*/60 * * * *", func() {
		if err := database.Exec("SELECT increase_salaries()").Error; err != nil {
			log.Fatalf("Failed to call stored procedure: %v", err)
		}

		fmt.Println("Salaries increased successfully!")
	})

	c.Start()
	select {}
}
