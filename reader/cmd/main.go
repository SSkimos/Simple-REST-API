package main

import (
	"github.com/curtrika/reader/pkg/common/db"
	"github.com/curtrika/reader/pkg/employees"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	employees.RegisterRoutes(r, h)

	r.Run(port)
}
