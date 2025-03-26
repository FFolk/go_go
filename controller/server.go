package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServer() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")

	dialactor := mysql.Open(dsn)
	db, err := gorm.Open(dialactor)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection successful")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	AuthController(router, db)
	CustomerController(router, db)
	router.Run()
}
