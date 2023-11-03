package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/baza-trainee/walking-school-backend/docs"
	"github.com/baza-trainee/walking-school-backend/internal/api"
	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/logger"
	"github.com/baza-trainee/walking-school-backend/internal/service"
	"github.com/baza-trainee/walking-school-backend/internal/storage"
	_ "github.com/swaggo/swag"
)

const timeoutLimit = 5

// @title Walking-School backend API
// @version 1.0
// tag.name "-----tag.name-----"
// tag.description "-----tag.description-----"
// @contact.name Yehor Tverytinov
// @contact.email etverya12@gmail.com
// @host localhost:7000
// host walking-school.site
// @BasePath /api/v1

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err.Error())

		return
	}

	log := logger.SetupLogger(cfg.LogLevel)
	log.Info("Server started")

	storage, err := storage.NewStorage(cfg.DB)
	if err != nil {
		log.Error("NewStorage error: ", err.Error())

		return
	}

	defer func() {
		if err := storage.DB.Client().Disconnect(context.TODO()); err != nil {
			log.Error("Storage Disconnect error: ", err.Error())

			return
		}
	}()

	service, err := service.NewService(storage, cfg)
	if err != nil {
		log.Error("New Service error: ", err.Error())

		return
	}

	server := api.NewServer(cfg, service, log)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.HTTPServer.Listen(cfg.Server.AppAddress); err != nil {
			log.Error("Start and Listen", "error", err.Error())
		}
	}()

	<-quit

	if err := server.HTTPServer.ShutdownWithTimeout(timeoutLimit * time.Second); err != nil {
		log.Error("ShutdownWithTimeout", "error", err.Error())
	}

	if err := server.HTTPServer.Shutdown(); err != nil {
		log.Error("Server shutdown", "error", err.Error())
	}

	log.Info("Server stopped")

}
