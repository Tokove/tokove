package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
	JWT struct {
		Secret      string
		ExpireHours int
	}
}

var AppConfig *Config

func InitConfig() {
	// 设置要读取的配置文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic("读取配置文件失败: " + err.Error())
	}

	// 解析
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		panic("解析配置文件失败: " + err.Error())
	}

	initDB()
	initRedis()
}
