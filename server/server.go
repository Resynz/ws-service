/**
 * @Author: Resynz
 * @Date: 2021/7/19 14:29
 */
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ws-service/common"
	"ws-service/config"
	"ws-service/controller"
	"ws-service/controller/api"
	"ws-service/middleware"
)

func StartServer() {
	gin.SetMode(config.Conf.Mode)
	app := gin.New()
	app.MaxMultipartMemory = 8 << 20 // 8mb
	app.Use(gin.Recovery())
	app.Use(middleware.Logger())

	app.GET("/ping", common.AuthDetection(controller.Ping))

	apiGroup := app.Group("/api")
	apiGroup.GET("/ws-url", common.AuthDetection(api.GetWsUrl))

	if err := app.Run(fmt.Sprintf(":%d", config.Conf.AppPort)); err != nil {
		log.Fatalf("start server failed! error:%s\n", err.Error())
	}
}
