package main

import (
	"fmt"
	"json-go/Jsonconverter"

	_ "json-go/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           JWT-Authentication API
// @version         1.0
// @description     This is eg of JWT Implementation.
// @termsOfService  http://demo.com

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8081

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/body", Jsonconverter.ValidateJson)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
