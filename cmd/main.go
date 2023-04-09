package main

import (
	"go_rocket_launch_sub/internal/app/server"
	"go_rocket_launch_sub/internal/config"
	"go_rocket_launch_sub/internal/pkg/handler"
	"go_rocket_launch_sub/internal/pkg/repository"
	"go_rocket_launch_sub/internal/pkg/service"
	"log"
	"os"
)

func main() {
	config.InitEnvs()
	config.InitDB()

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(server.App)
	if error := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); error != nil {
		log.Fatal("Error occured %s", error.Error())
	}
}
