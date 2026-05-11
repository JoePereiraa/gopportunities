package main

import (
	"github.com/JoePereiraa/gopportunities/config"
	"github.com/JoePereiraa/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	//Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
