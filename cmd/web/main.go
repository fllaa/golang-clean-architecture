package main

import (
	"fmt"
	"golang-clean-architecture/internal/config"

	_ "golang-clean-architecture/docs"
)

// @title           Golang Clean Architecture
// @version         1.0.0
// @description     Golang Clean Architecture

// @host      localhost:3000
// @BasePath  /api/

// @tag.name Address
// @tag.description Address Management API

// @tag.name Contact
// @tag.description Contact Management API

// @tag.name User
// @tag.description User Management API
func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)
	producer := config.NewKafkaProducer(viperConfig, log)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		Producer: producer,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
