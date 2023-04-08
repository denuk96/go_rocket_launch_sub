package main

import (
	"go_rocket_launch_sub/internal/app/server"
	"go_rocket_launch_sub/internal/pkg/handler"
	"go_rocket_launch_sub/internal/pkg/repository"
	"go_rocket_launch_sub/internal/pkg/service"
	"log"
)

func main() {
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(server.App)
	if error := server.Run("3000", handlers.InitRoutes()); error != nil {
		log.Fatal("Error occured %s", error.Error())
	}
}
