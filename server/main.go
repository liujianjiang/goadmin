package main

import (
	"log"

	"github.com/liujianjiang/goadmin/server/core"
	"github.com/liujianjiang/goadmin/server/global"
	"github.com/liujianjiang/goadmin/server/initialize"
	"github.com/liujianjiang/goadmin/server/utils/antspool"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	//初始化协程池
	if err := antspool.InitPools(); err != nil {
		log.Printf("common.InitPools err: %v", err)
	}
	core.RunWindowsServer()
}
