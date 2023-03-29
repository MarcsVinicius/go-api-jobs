package main

import (
	"github.com/marcsvinicius/go-api-jobs/config"
	"github.com/marcsvinicius/go-api-jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Inicializa Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Inicializa Rotas
	router.Initialize()
}
