package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"retrieval/pkg/common/db"
	"retrieval/pkg/employees"
	"time"
)

func main() {
	time.Sleep(8 * time.Second)
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	employees.RegisterRoutes(r, h)

	r.Run(port)
}
