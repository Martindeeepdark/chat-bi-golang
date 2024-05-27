// main.go
package main

import (
	"chat-bi-golang/configs"
	"chat-bi-golang/route"
	"chat-bi-golang/services/sessionmanager"
	"fmt"
)

func main() {
	// 加载配置
	config, err := configs.LoadConfig("config.yml")
	if err != nil {
		fmt.Println("加载配置失败:", err)
		return
	}

	// 初始化 WebSocket 会话
	sessionmanager.InitSession(config.XfApi.AppId, config.XfApi.ApiKey, config.XfApi.ApiSecret, config.XfApi.HostUrl)

	// 获取路由引擎，传递配置
	router := route.Setroute(config)

	// 启动 Gin 引擎并监听端口
	err = router.Run(":8080")
	if err != nil {
		fmt.Println("启动 Gin 引擎失败:", err)
		return
	}

	fmt.Println("Gin 引擎已启动，监听端口:", ":8080")
}
