package main

import (
	"fmt"
	"json-go/Jsonconverter"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/body", Jsonconverter.ValidateJson)

	router.SetTrustedProxies(nil)
	// host := flag.String("host", "", "MyIP")
	// port := flag.String("port", "", "port")
	// flag.Parse()
	// // router.Run("localhost:8080")
	// MyAddr := *host + *port
	vi := viper.New()
	vi.SetConfigFile("config.yaml")
	vi.ReadInConfig()
	IP := vi.GetString("info.ServerIP")
	Port := vi.GetString("info.Port")
	MyAddr := fmt.Sprint(IP + ":" + Port)
	router.Run(MyAddr)
}
