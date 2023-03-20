package main

import (
	"go_rocket_launch_sub/internal/app/server"
	"go_rocket_launch_sub/internal/pkg/handler"
	"log"
)

func main() {
	server := new(server.App)
	handler := new(handler.Handler)

	if error := server.Run("3000", handler.InitRoutes()); error != nil {
		log.Fatal("Error occured %s", error.Error())
	}
}
