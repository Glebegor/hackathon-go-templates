package main

import (
	"fmt"

	"github.com/Glebegor/hackathon-go-templates/bootstrap"
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

	// Config init
	server := bootstrap.NewServer(
		fmt.Sprintf("%s:%d", config.Server.HOST, config.Server.PORT),
		nil, // TODO: replace with actual router
	)

	// Create app
	app := bootstrap.NewApp(config, server, logger)

	app.Run()

}
