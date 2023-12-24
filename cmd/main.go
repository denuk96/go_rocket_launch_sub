package main

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go_rocket_launch_sub/internal/app/server"
	"go_rocket_launch_sub/internal/config"
	"go_rocket_launch_sub/internal/pkg/handler"
	"go_rocket_launch_sub/internal/pkg/repository"
	"go_rocket_launch_sub/internal/pkg/service"
	"gorm.io/gorm"
	"net/http"
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

	srv := new(server.App)

	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info("HTTP server stopped")
			} else {
				log.Fatalf("Error occurred while running HTTP server: %s", err.Error())
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}
	closeDbConnection(db)

	log.Println("Server is shutting down")
}

func closeDbConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error on extracting sql.DB from GORM: %s", err)
	} else {
		err = sqlDB.Close()
		if err != nil {
			log.Printf("Error on closing the database: %s", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
