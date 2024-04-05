package main

import (
	"github.com/curtrika/publisher/pkg/common/db"
	"github.com/curtrika/publisher/pkg/employees"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	NatsUrl := viper.Get("NATS_URL").(string)

	r := gin.Default()
	h := db.Init(NatsUrl)

	employees.RegisterRoutes(r, h)

	r.Run(port)
}
