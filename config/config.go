package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config 全局配置
type Config struct {
	ServerPort string
	DBPath     string
}

// AppConfig 全局配置实例
var AppConfig Config

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigFile("config.yaml") // 配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read configuration file:", err)
	}

	// 读取配置信息
	AppConfig.ServerPort = viper.GetString("server.port")
	AppConfig.DBPath = viper.GetString("database.path")
}
