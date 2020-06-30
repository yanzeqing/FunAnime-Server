package main

import (
	"github.com/yanzeqing/FunAnime-Server/cache"
	"github.com/yanzeqing/FunAnime-Server/model"
	"github.com/yanzeqing/FunAnime-Server/util/conf"
	"github.com/yanzeqing/FunAnime-Server/util/logger"
)

func initHandler(runType string) {
	logger.Init()
	conf.Init(runType)
	model.DatabaseInit()
	cache.Redis()
}
