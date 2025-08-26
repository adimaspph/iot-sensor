package main

import (
	"fmt"
	"iot-sensor/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	validate := config.NewValidator(viperConfig)
	app := config.NewEcho(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	port := viperConfig.GetInt("APP_PORT")

	log.Infof("Starting server on port %d", port)
	err := app.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
