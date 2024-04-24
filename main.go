package main

import (
	"github.com/AndreD23/go-api/config"
	"github.com/AndreD23/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.BuildAndRun()
}
