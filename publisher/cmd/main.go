package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"publisher/pkg/common/utils"
	"publisher/pkg/employees"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	NatsUrl := viper.Get("NATS_URL").(string)

	r := gin.Default()
	nc := utils.InitNatsConnection(NatsUrl)
	defer utils.CloseConnection(nc)

	employees.RegisterRoutes(r, nc)

	r.Run(port)
}
