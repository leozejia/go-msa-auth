package main

import (
	"go-msa-auth/config"
	"go-msa-auth/server"
	"log"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化服务器
	if err := server.InitServer(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
