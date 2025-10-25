package main

import (
	"backend/global"
	"backend/model"
	"log"

	"backend/config"
	"backend/router"
)

func main() {
	config.InitConfig()
	if err := global.Db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	
	r := router.SetupRouter()

	r.Run(config.AppConfig.App.Port)
}
