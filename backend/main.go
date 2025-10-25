package main

import (
	"backend/config"
	"backend/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	r.Run(config.AppConfig.App.Port)
}
