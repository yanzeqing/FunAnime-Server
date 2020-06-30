package main

import (
	"github.com/yanzeqing/FunAnime-Server/router"
	barrage "github.com/yanzeqing/FunAnime-Server/service/websocket"
	"github.com/yanzeqing/FunAnime-Server/util/logger"
)

func main() {
	initHandler("dev")
	// websocket服务 监听8090
	go barrage.Main()
	err := router.NewRouter().Run(":8088")
	if err != nil {
		logger.Fatal("start_serve_failed", logger.Fields{"err": err})
		return
	}
}
