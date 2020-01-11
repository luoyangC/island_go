package main

import (
	"island/router"
	"island/settings"
)

// @title Island Go API
// @version 1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
// @host 127.0.0.1:3000
func main() {
	// 从配置文件读取配置
	settings.Init()
	// 装载路由
	r := router.NewRouter()
	// 启动服务器
	r.Run(":3000")
}
