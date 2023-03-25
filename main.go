package main

import (
	"github.com/gin-gonic/gin"
	"json-go/Jsonconverter"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/body",Jsonconverter.ValidateJson)

	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
