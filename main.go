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
/*
{
    "GlobalEnbId":{
        "mcc":208,
        "mnc":93,
        "enbid":1
    },
    "SupportedTAList":[
        {
            "TAC":12580,
            "BroadCastPLMN":[
                {
                "MCC":208,
                "MNC":93,
                "b":[
                    "i",
                    "j"
                ],
            "SupportedSliceList":[
                {
                    "SST":1,
                    "SD":234
                }
            ]
        }
            ]
        }
    ],
            "PagingDRx":"v23"
}

*/