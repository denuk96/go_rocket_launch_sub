package main

import (
	"go_rocket_launch_sub/internal/app/server"
)

func main() {
	server := new(server.App)
	server.Run("3000")
}
