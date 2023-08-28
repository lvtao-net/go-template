package config

import (
	"github.com/joho/godotenv"
)

func InitConfig() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		return
	}
}
