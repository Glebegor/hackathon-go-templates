package main

import (
	"fmt"

	"github.com/Glebegor/hackathon-go-templates/api/handlers"
	"github.com/Glebegor/hackathon-go-templates/api/routers"
	"github.com/Glebegor/hackathon-go-templates/bootstrap"
	repositories "github.com/Glebegor/hackathon-go-templates/repositories"
	usecase "github.com/Glebegor/hackathon-go-templates/usecases"
)

func main() {
	// Load configuration
	config, err := bootstrap.NewConfig("./config/config.json", "dev")
	if err != nil {
		panic(err)
	}

	// Logger setup
	logger := bootstrap.NewLogger(bootstrap.LoggerConfig{
		Env:   config.Server.ENV,
		Level: "debug",
	})
	logger.Info("starting service",
		"host", config.Server.HOST,
		"port", config.Server.PORT,
		"name", config.Server.NAME,
		"env", config.Server.ENV,
	)

	// DB setup
	db, err := bootstrap.NewGormDB(*&config.DbConfig, logger)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		panic(err)
	}

	// Dependencies setup
	repos := repositories.NewRepositories(db, config, logger)
	usecases := usecase.NewUsecases(config, logger, repos)
	handlers := handlers.NewHandlers(config, logger, usecases)

	// Router setup
	router := routers.NewRouter(handlers, config.Server.ENV)

	// Config init
	server := bootstrap.NewServer(
		fmt.Sprintf("%s:%d", config.Server.HOST, config.Server.PORT),
		router,
	)

	// Create app
	app := bootstrap.NewApp(config, server, logger)

	app.Run()

}
