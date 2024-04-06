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
	nc, err := utils.InitNATSConnection(NatsUrl)
	defer utils.CloseConnection(nc)

	if err != nil {
		//TODO: тык
	}

	employees.RegisterRoutes(r, nc)

	r.Run(port)
}
