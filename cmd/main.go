package main

import (
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/app/server"
	"go_rocket_launch_sub/internal/config"
	"go_rocket_launch_sub/internal/pkg/handler"
	"go_rocket_launch_sub/internal/pkg/repository"
	"go_rocket_launch_sub/internal/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitEnvs()
	db := config.InitDB()
	smtpCreds := config.InitSmtp()

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories, smtpCreds)
	handlers := handler.NewHandler(services)

	app := new(server.App)

	go func() {
		if err := app.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occurred while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server is shutting down")
}
